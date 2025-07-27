package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"amocrm2.0/internal/config"
	"amocrm2.0/internal/core/amocrm"
	"amocrm2.0/internal/infrastructure/transport/http/dto"
	"amocrm2.0/internal/usecases"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

const (
	authGrantType = "authorization_code"
	contentType   = "application/json"
)

type AccountHandlers struct {
	AccountUC *usecases.AccountUC
	Cfg       *config.Config
}

func newAccountHandlers(uc *usecases.AccountUC, cfg *config.Config) *AccountHandlers {
	return &AccountHandlers{
		AccountUC: uc,
		Cfg:       cfg,
	}
}

// GET method
func (h *AccountHandlers) HandleAuth(w http.ResponseWriter, r *http.Request) {
	integrationInfo := getIntegrationInfoFromQuery(r)

	account, err := exchangeTokens(integrationInfo, &h.Cfg.Integration)
	if err != nil {
		logrus.Error(err)
		sendErrorResponse(w, r, http.StatusBadRequest, err)
		return
	}

	if err := h.AccountUC.Add(account); err != nil {
		logrus.Error(err)
		sendErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	sendOKResponse(w, r, http.StatusCreated, account, "successful auth")
}

func exchangeTokens(integration *dto.IntegrationInfoRequest, cfg *config.Integration) (*amocrm.Account, error) {
	req := dto.NewExchangeTokensRequest(
		integration.AuthCode,
		cfg.ClientID,
		cfg.SecrestKey,
		cfg.RedirectURL,
		authGrantType,
	)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal json: %v", err)
	}

	exchangeTokensURL := makeExchangeTokensURL(integration.Referer)

	resp, err := http.Post(
		exchangeTokensURL,
		contentType,
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to exchange tokens: status code %d", resp.StatusCode)
	}

	var exchangeResponse dto.ExchangeTokensResponse
	if err := render.DecodeJSON(resp.Body, &exchangeResponse); err != nil {
		return nil, fmt.Errorf("failed to decode exchange response: %v", err)
	}

	account := exchangeResponse.ToDomainAccount()
	return &account, nil
}

func (h *AccountHandlers) Add(w http.ResponseWriter, r *http.Request) {

}

func (h *AccountHandlers) GetByID(w http.ResponseWriter, r *http.Request) {

}

func (h *AccountHandlers) GetAll(w http.ResponseWriter, r *http.Request) {

}

func (h *AccountHandlers) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *AccountHandlers) Delete(w http.ResponseWriter, r *http.Request) {

}

func getIntegrationInfoFromQuery(r *http.Request) *dto.IntegrationInfoRequest {
	integrationInfo := &dto.IntegrationInfoRequest{
		AuthCode: r.URL.Query().Get("code"),
		Referer:  r.URL.Query().Get("referer"),
	}
	return integrationInfo
}

func makeExchangeTokensURL(referer string) string {
	return fmt.Sprintf("https://%s/oauth2/access_token", referer)
}
