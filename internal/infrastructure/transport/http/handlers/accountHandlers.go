package handlers

import (
	"net/http"

	"amocrm2.0/internal/usecases"
)

type AccountHandlers struct {
	AccountUC *usecases.AccountUC
}

func newAccountHandlers(uc *usecases.AccountUC) *AccountHandlers {
	return &AccountHandlers{
		AccountUC: uc,
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
