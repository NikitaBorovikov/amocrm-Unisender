package dto

import (
	"net/mail"
	"time"

	"amocrm2.0/internal/core/amocrm"
)

type ErrorResponse struct {
	Msg string `json:"error"`
}

type OKResponse struct {
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

type ExchangeTokensResponse struct {
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Expires      int64  `json:"expires_in"`
}

type RefreshAccessTokenResponse struct {
	TokenType    string `json:"token_type"`
	Expires      int64  `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type GetAccountInfoResponse struct {
	ID int `json:"id"`
}

type ContactAmoCRMResponse struct {
	Embedded Embedded `json:"_embedded"`
}

type Embedded struct {
	Contacts []Contact `json:"contacts"`
}

type Contact struct {
	ID           int            `json:"id"`
	Name         string         `json:"name"`
	CustomFields []CustomFields `json:"custom_fields_values"`
	AccountID    int            `json:"account_id"`
}

type CustomFields struct {
	FieldCode string   `json:"field_code"`
	Values    []Values `json:"values"`
}

type Values struct {
	Value interface{} `json:"value"`
}

func NewErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{
		Msg: err.Error(),
	}
}

func NewOKReponse(data interface{}, msg string) *OKResponse {
	return &OKResponse{
		Msg:  msg,
		Data: data,
	}
}

func (r *ExchangeTokensResponse) ToDomainAccount(domain string) amocrm.Account {
	return amocrm.Account{
		AccessToken:  r.AccessToken,
		RefreshToken: r.RefreshToken,
		Expires:      r.Expires,
		Domain:       domain,
		IssuedAt:     time.Now(),
	}
}

func (r *ContactAmoCRMResponse) ToDomainContacts() []amocrm.Contact {
	var contacts []amocrm.Contact
	for _, c := range r.Embedded.Contacts {
		contact := amocrm.Contact{
			ContactID: c.ID,
			AccountID: c.AccountID,
			Name:      c.Name,
		}

		hasValidEmail := false

		for _, field := range c.CustomFields {
			if field.FieldCode == "EMAIL" {
				if email, ok := field.Values[0].Value.(string); ok && isValidEmail(email) {
					contact.Email = email
					hasValidEmail = true
				}
			}
		}

		if hasValidEmail {
			contacts = append(contacts, contact)
		}
	}
	return contacts
}

func isValidEmail(email string) bool {
	if len(email) < 7 {
		return false
	}
	_, err := mail.ParseAddress(email)
	return err == nil
}
