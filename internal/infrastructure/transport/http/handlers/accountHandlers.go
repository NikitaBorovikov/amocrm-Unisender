package handlers

import (
	"net/http"
	"strconv"

	"amocrm2.0/internal/config"
	"amocrm2.0/internal/usecases"
	"github.com/sirupsen/logrus"
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

	if err := h.AccountHandlers.AccountUC.UpdateUnisenderKey(accountID, apiKey); err != nil {
		sendErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	sendOKResponse(w, r, http.StatusOK, nil, "Unisender key is updated")

	// Вызов функции первичной синхронизации
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
