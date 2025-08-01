package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"amocrm2.0/internal/core/amocrm"
	"amocrm2.0/internal/infrastructure/transport/http/dto"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

const (
	importContactsURL = "https://api.unisender.com/ru/api/importContacts?format=json"
)

func (h *Handlers) HandleFirstSync(accountID int) {
	contacts, err := h.GetContacts(accountID)
	if err != nil {
		logrus.Error(err)
		return
	}

	apiKey, err := h.UseCases.AccountUC.GetUnisenderKey(accountID)
	if err != nil {
		logrus.Errorf("failed to get Unisender API key: %v", err)
		return
	}

	importData := prepareUnisenderImportData(apiKey, contacts)
	if err := h.SendContactsToUnisender(importData); err != nil {
		logrus.Error(err)
		return
	}
	// TODO: Save contact in DB
}

func (h *Handlers) GetContacts(accountID int) ([]amocrm.Contact, error) {
	account, err := h.UseCases.AccountUC.GetByID(accountID)
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
	return contacts, nil
}

func (h *Handlers) SendContactsToUnisender(data *dto.UnisenderImportRequest) error {
	resp, err := sendImportContactsRequest(data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var unisenderResponse dto.UnisenderImportResponse
	if err := render.DecodeJSON(resp.Body, &unisenderResponse); err != nil {
		return fmt.Errorf("failed to decode JSON: %v", err)
	}
	return nil
}

func makeGetContactsURL(domain string) string {
	return fmt.Sprintf("https://%s/api/v4/contacts", domain)
}

func prepareUnisenderImportData(apiKey string, contacts []amocrm.Contact) *dto.UnisenderImportRequest {
	fieldsName := []string{"name", "email"}
	data := make([][]string, len(contacts))
	for idx, contact := range contacts {
		data[idx] = []string{contact.Name, contact.Email}
	}

	return &dto.UnisenderImportRequest{
		APIKey:     apiKey,
		FieldNames: fieldsName,
		Data:       data,
	}
}

func sendImportContactsRequest(data *dto.UnisenderImportRequest) (*http.Response, error) {
	req, err := prepareSendContactsRequest(data)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare send contacts request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unisender error: status code %d", resp.StatusCode)
	}
	return resp, nil
}

func prepareSendContactsRequest(data *dto.UnisenderImportRequest) (*http.Request, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		importContactsURL,
		bytes.NewBuffer(jsonData),
	)
	return req, err
}
