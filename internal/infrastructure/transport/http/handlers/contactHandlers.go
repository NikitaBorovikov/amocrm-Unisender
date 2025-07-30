package handlers

import (
	"fmt"
	"net/http"

	"amocrm2.0/internal/core/amocrm"
	"amocrm2.0/internal/infrastructure/transport/http/dto"
	"amocrm2.0/internal/usecases"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

type ContactHandlers struct {
	ContactUC *usecases.ContactUC
	AccountUC *usecases.AccountUC
}

func newContactHandlers(contactUC *usecases.ContactUC, accountUC *usecases.AccountUC) *ContactHandlers {
	return &ContactHandlers{
		ContactUC: contactUC,
		AccountUC: accountUC,
	}
}

func (h *Handlers) GetContacts(accountID int) ([]amocrm.Contact, error) {
	account, err := h.AccountHandlers.AccountUC.GetByID(accountID)
	if err != nil {
		return nil, fmt.Errorf("failed to get account info: %v", err)
	}

	if !isValidAccessToken(account) {
		if err := h.refreshAccessToken(account.AccountID); err != nil {
			return nil, fmt.Errorf("failed to refresh token: %v", err)
		}
	}

	getContactURl := makeGetContactsURL(account.Domain)
	resp, err := sendGetRequestToAmoCRM(account.AccessToken, getContactURl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var amoCRMResponse dto.ContactAmoCRMResponse
	if err := render.DecodeJSON(resp.Body, &amoCRMResponse); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %v", err)
	}

	contacts := amoCRMResponse.ToDomainContacts()
	logrus.Info(contacts) //REMOVE: for testing
	return contacts, nil
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

func makeGetContactsURL(domain string) string {
	return fmt.Sprintf("https://%s/api/v4/contacts", domain)
}
