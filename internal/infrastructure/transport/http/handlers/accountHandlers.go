package handlers

import (
	"net/http"

	"amocrm2.0/internal/config"
	"amocrm2.0/internal/usecases"
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
