package handlers

import (
	"fmt"
	"net/http"
	"net/mail"
	"net/url"
	"strconv"

	"amocrm2.0/internal/core/amocrm"
	"github.com/sirupsen/logrus"
)

const (
	maxAmountCustomFields       = 20
	maxAmountContactsInWebbhook = 10
	addEventType                = "add"
	updateEventType             = "update"
)

func (h *Handlers) ReceiveContactWebhook(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		logrus.Errorf("failed to parse form: %v", err)
		return
	}
	accountID := r.PostForm.Get("account[id]")
	logrus.Infof("ACCOUNT_ID: %s", accountID) // REMOVE: for testing
	addedContacts := parseContactsFromWebhook(r.PostForm, addEventType)
	updatedContacts := parseContactsFromWebhook(r.PostForm, updateEventType)

	// Может тут записывать в очередь не по одному контакту, а сразу список
	for _, contact := range addedContacts {
		h.processContact(contact, addEventType)
	}

	for _, contact := range updatedContacts {
		h.processContact(contact, updateEventType)
	}
}

func (h *Handlers) processContact(contact amocrm.Contact, eventType string) {
	// make a task
	// add in a queue
}

func parseContactsFromWebhook(form url.Values, eventType string) []amocrm.Contact {
	var contacts []amocrm.Contact

	for i := 0; i < maxAmountContactsInWebbhook; i++ {
		idStr := form.Get(fmt.Sprintf("contacts[%s][%d][id]", eventType, i))
		if idStr == "" {
			break // контакты закончились
		}

		accountID := form.Get(fmt.Sprintf("contacts[%s][%d][account_id]", eventType, i))
		name := form.Get(fmt.Sprintf("contacts[%s][%d][name]", eventType, i))

		contact := amocrm.Contact{
			ContactID: atoi(idStr),
			AccountID: atoi(accountID),
			Name:      name,
		}

		email := findEmailInForm(form, i, eventType)
		if !isValidEmail(email) {
			continue
		}
		contact.Email = email

		contacts = append(contacts, contact)
	}
	return contacts
}

func atoi(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		logrus.Errorf("failed to convert to int: %v", err)
		return 0
	}
	return val
}

func findEmailInForm(form url.Values, contactID int, eventType string) string {
	for i := 0; i < maxAmountCustomFields; i++ {
		codeKey := fmt.Sprintf("contacts[%s][%d][custom_fields][%d][code]", eventType, contactID, i)
		if form.Get(codeKey) == "EMAIL" {
			return form.Get(fmt.Sprintf("contacts[%s][%d][custom_fields][%d][values][0][value]", eventType, contactID, i))
		}
	}
	return ""
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
