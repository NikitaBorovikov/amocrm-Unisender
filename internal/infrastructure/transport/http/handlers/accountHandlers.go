package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"amocrm2.0/internal/infrastructure/transport/http/dto"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

func (h *Handlers) ReceiveUnisenderKey(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		logrus.Error(err)
		sendErrorResponse(w, r, http.StatusBadRequest, err)
		return
	}

	apiKey := r.FormValue("unisender_key")
	accountIDStr := r.FormValue("account_id")

	accountID, err := strconv.Atoi(accountIDStr)
	if err != nil {
		logrus.Error(err)
		sendErrorResponse(w, r, http.StatusBadRequest, err)
		return
	}
	logrus.Infof("accountID: %d key: %s", accountID, apiKey)

	if err := h.UseCases.AccountUC.UpdateUnisenderKey(accountID, apiKey); err != nil {
		sendErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	sendOKResponse(w, r, http.StatusOK, nil, "Unisender key is updated")

	// Вызов функции первичной синхронизации
}

func (h *Handlers) GetAccountID(accessToken, domain string) (int, error) {
	getAccountURL := makeGetAccountURL(domain)

	resp, err := sendGetRequestToAmoCRM(accessToken, getAccountURL)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var accountInfo dto.GetAccountInfoResponse
	if err := render.DecodeJSON(resp.Body, &accountInfo); err != nil {
		return 0, fmt.Errorf("failed to decode JSON: %v", err)
	}
	return accountInfo.ID, nil
}

func makeGetAccountURL(domain string) string {
	return fmt.Sprintf("https://%s/api/v4/account", domain)
}
