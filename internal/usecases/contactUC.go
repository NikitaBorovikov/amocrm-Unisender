package usecases

import "amocrm2.0/internal/core/amocrm"

type ContactUC struct {
	ContactRepo amocrm.ContactRepo
}

func NewContactUC(contactRepo amocrm.ContactRepo) *ContactUC {
	return &ContactUC{
		ContactRepo: contactRepo,
	}
}

func (uc *ContactUC) Add(contact *amocrm.Contact) error {
	return nil
}

func (uc *ContactUC) GetByID(contactID int) (*amocrm.Contact, error) {
	return nil, nil
}

func (uc *ContactUC) GetAll() ([]amocrm.Contact, error) {
	return nil, nil
}

func (uc *ContactUC) Update(contact *amocrm.Contact) error {
	return nil
}

func (uc *ContactUC) Delete(contactID int) error {
	return nil
}
