package handlers

import (
	"net/http"

	"amocrm2.0/internal/usecases"
)

type ContactHandlers struct {
	ContactUC *usecases.ContactUC
}

func newContactHandlers(uc *usecases.ContactUC) *ContactHandlers {
	return &ContactHandlers{
		ContactUC: uc,
	}
}

func (h *ContactHandlers) Add(w http.ResponseWriter, r *http.Request) {

}

func (h *ContactHandlers) GetByID(w http.ResponseWriter, r *http.Request) {

}

func (h *ContactHandlers) GetAll(w http.ResponseWriter, r *http.Request) {

}

func (h *ContactHandlers) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *ContactHandlers) Delete(w http.ResponseWriter, r *http.Request) {

}
