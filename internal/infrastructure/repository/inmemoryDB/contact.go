package inmemorydb

import (
	"errors"
	"sync"

	"amocrm2.0/internal/core/amocrm"
)

var (
	errContactNotExists = errors.New("such contactID does not exist")
)

type ContactRepo struct {
	contacts map[int]amocrm.Contact
	mu       sync.RWMutex
}

func NewContactRepo() *ContactRepo {
	return &ContactRepo{
		contacts: make(map[int]amocrm.Contact),
	}
}

func (r *ContactRepo) Add(contact *amocrm.Contact) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.contacts[contact.ContactID] = *contact
	return nil
}

func (r *ContactRepo) GetByID(conatctID int) (*amocrm.Contact, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	val, ok := r.contacts[conatctID]
	if !ok {
		return nil, errContactNotExists
	}
	return &val, nil
}

func (r *ContactRepo) GetAll() ([]amocrm.Contact, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	contacts := make([]amocrm.Contact, 0, len(r.contacts))
	for _, val := range r.contacts {
		contacts = append(contacts, val)
	}
	return contacts, nil
}

func (r *ContactRepo) Update(contact *amocrm.Contact) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.contacts[contact.ContactID]
	if !ok {
		return errContactNotExists
	}

	r.contacts[contact.ContactID] = *contact
	return nil
}

func (r *ContactRepo) Delete(contactID int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.contacts[contactID]
	if !ok {
		return errContactNotExists
	}

	delete(r.contacts, contactID)
	return nil
}
