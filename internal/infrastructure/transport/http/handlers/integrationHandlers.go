package handlers

import (
	"net/http"

	"amocrm2.0/internal/usecases"
)

type IntegrationHandlers struct {
	IntegrationUC *usecases.IntegrationUC
}

func newIntegrationHandlers(uc *usecases.IntegrationUC) *IntegrationHandlers {
	return &IntegrationHandlers{
		IntegrationUC: uc,
	}
}

func (h *IntegrationHandlers) Add(w http.ResponseWriter, r *http.Request) {

}

func (h *IntegrationHandlers) GetByID(w http.ResponseWriter, r *http.Request) {

}

func (h *IntegrationHandlers) GetAll(w http.ResponseWriter, r *http.Request) {

}

func (h *IntegrationHandlers) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *IntegrationHandlers) Delete(w http.ResponseWriter, r *http.Request) {

}
