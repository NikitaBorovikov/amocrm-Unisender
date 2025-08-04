package handlers

import (
	"fmt"
	"net/http"
	"net/mail"
	"net/url"
	"strconv"

	"amocrm2.0/internal/core/amocrm"
	"amocrm2.0/internal/infrastructure/queue"
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
	accountIDStr := r.PostForm.Get("account[id]")
	accountID, err := strconv.Atoi(accountIDStr)
	if err != nil {
		logrus.Errorf("failed to convent accountID to int: %v", err)
		return
	}

	addedContacts := parseContactsFromWebhook(r.PostForm, addEventType)
	updatedContacts := parseContactsFromWebhook(r.PostForm, updateEventType)

	h.processContact(addedContacts, addEventType, accountID)
	h.processContact(updatedContacts, updateEventType, accountID)
}

func (h *Handlers) processContact(contacts []amocrm.Contact, eventType string, accountID int) {
	// Если нет валидных котнактов, то нет смысла записывать задачу в очередь
	if len(contacts) == 0 {
		return
	}

	task := queue.SyncContactsTask{
		AccountID: accountID,
		EventType: eventType,
		TaskType:  "webhook_sync",
		Contacts:  contacts,
	}

	_, err := h.Producer.AddSyncContactsTask(task)
	if err != nil {
		logrus.Errorf("failed to queue a task: %v", err)
		return
	}
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
