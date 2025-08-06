package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"amocrm2.0/internal/core/amocrm"
	"amocrm2.0/internal/infrastructure/transport/http/dto"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

const (
	authGrantType    = "authorization_code"
	refreshGrantType = "refresh_token"
	contentType      = "application/json"
)

func (h *Handlers) HandleAuth(w http.ResponseWriter, r *http.Request) {
	integrationInfo := getIntegrationInfoFromQuery(r)

	account, err := h.exchangeTokens(integrationInfo)
	if err != nil {
		logrus.Error(err)
		sendErrorResponse(w, r, http.StatusBadRequest, err)
		return
	}

	// Нужно получить ID аккаунта amoCRM, прежде чем записать в БД
	accountID, err := h.GetAccountID(account.AccessToken, account.Domain)
	if err != nil {
		logrus.Error(err)
		sendErrorResponse(w, r, http.StatusBadRequest, err)
		return
	}
	account.AccountID = accountID

	if err := h.UseCases.AccountUC.Add(account); err != nil {
		logrus.Error(err)
		sendErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	sendOKResponse(w, r, http.StatusCreated, account, "successful auth")
}

func (h *Handlers) exchangeTokens(integration *dto.IntegrationInfoRequest) (*amocrm.Account, error) {
	req := dto.NewExchangeTokensRequest(
		integration.AuthCode,
		h.Cfg.Integration.ClientID,
		h.Cfg.Integration.SecrestKey,
		h.Cfg.Integration.RedirectURL,
		authGrantType,
	)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal json: %v", err)
	}

	exchangeTokensURL := makeAuthURL(integration.Domain)

	resp, err := sendAuthRequest(reqBody, exchangeTokensURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var exchangeResponse dto.ExchangeTokensResponse
	if err := render.DecodeJSON(resp.Body, &exchangeResponse); err != nil {
		return nil, fmt.Errorf("failed to decode exchange response: %v", err)
	}

	account := exchangeResponse.ToDomainAccount(integration.Domain)
	return &account, nil
}

func (h *Handlers) refreshAccessToken(accountID int) error {
	account, err := h.UseCases.AccountUC.GetByID(accountID)
	if err != nil {
		return fmt.Errorf("failed to get account info: %v", err)
	}

	req := dto.NewRefreshAccessTokenRequest(
		h.Cfg.Integration.ClientID,
		h.Cfg.Integration.SecrestKey,
		h.Cfg.Integration.RedirectURL,
		account.Domain,
		refreshGrantType,
	)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal json: %v", err)
	}

	refreshURL := makeAuthURL(account.Domain)

	resp, err := sendAuthRequest(reqBody, refreshURL)
	if err != nil {
		logrus.Error(err)
		return err
	}
	defer resp.Body.Close()

	var refreshResponse dto.RefreshAccessTokenResponse
	if err := render.DecodeJSON(resp.Body, &refreshResponse); err != nil {
		return fmt.Errorf("failed to decode refresh response: %v", err)
	}

	account.AccessToken = refreshResponse.AccessToken
	account.RefreshToken = refreshResponse.RefreshToken
	account.Expires = refreshResponse.Expires
	account.IssuedAt = time.Now()

	if err := h.UseCases.AccountUC.Update(account); err != nil {
		return fmt.Errorf("failed to update tokens in DB: %v", err)
	}
	return nil
}

func getIntegrationInfoFromQuery(r *http.Request) *dto.IntegrationInfoRequest {
	integrationInfo := &dto.IntegrationInfoRequest{
		AuthCode: r.URL.Query().Get("code"),
		Domain:   r.URL.Query().Get("referer"),
	}
	return integrationInfo
}

func makeAuthURL(domain string) string {
	return fmt.Sprintf("https://%s/oauth2/access_token", domain)
}

func sendAuthRequest(reqBody []byte, url string) (*http.Response, error) {
	resp, err := http.Post(
		url,
		contentType,
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("amoCRM error: status code %d", resp.StatusCode)
	}
	return resp, nil
}

func isValidAccessToken(account *amocrm.Account) bool {
	expiresAt := account.IssuedAt.Add(time.Duration(account.Expires) * time.Second)
	now := time.Now()
	return now.Before(expiresAt)
}
