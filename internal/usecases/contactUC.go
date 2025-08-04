package usecases

import (
	"strings"

	"amocrm2.0/internal/core/amocrm"
	"github.com/sirupsen/logrus"
)

const (
	errCodeDublicate = "1062"
)

type ContactUC struct {
	ContactRepo amocrm.ContactRepo
}

func NewContactUC(contactRepo amocrm.ContactRepo) *ContactUC {
	return &ContactUC{
		ContactRepo: contactRepo,
	}
}

func (uc *ContactUC) HandleSaveNewContactData(contacts []amocrm.Contact, syncStatus bool, eventType string) {
	switch eventType {
	case "add":
		uc.handleAddContacts(contacts, syncStatus)
	case "update":
		uc.handleUpdateContacts(contacts, syncStatus)
	default:
		logrus.Info("invalid event type")
	}
}

func (uc *ContactUC) handleAddContacts(contacts []amocrm.Contact, syncStatus bool) {
	for _, contact := range contacts {
		contact.SyncStatus = syncStatus
		if err := uc.Add(&contact); err != nil {
			if strings.Contains(err.Error(), errCodeDublicate) { // Если такой контакт уже есть, то обновляем его данные
				uc.Update(&contact)
				continue
			}
			logrus.Errorf("failed to add contact (id = %d): %v", contact.ContactID, err)
		}
	}
}

func (uc *ContactUC) handleUpdateContacts(contacts []amocrm.Contact, syncStatus bool) {
	for _, contact := range contacts {
		contact.SyncStatus = syncStatus
		if err := uc.Update(&contact); err != nil {
			logrus.Errorf("failed to update contact (id = %d): %v", contact.ContactID, err)
		}
	}
}

func (uc *ContactUC) Add(contact *amocrm.Contact) error {
	return uc.ContactRepo.Add(contact)
}

func (uc *ContactUC) Update(contact *amocrm.Contact) error {
	return uc.ContactRepo.Update(contact)
}
