package mysqldb

import (
	"sync"

	"amocrm2.0/internal/core/amocrm"
	"gorm.io/gorm"
)

type ContactRepo struct {
	db *gorm.DB
	mu sync.RWMutex
}

func NewContactRepo(db *gorm.DB) *ContactRepo {
	return &ContactRepo{db: db}
}

func (r *ContactRepo) Add(contact *amocrm.Contact) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	return nil
}

func (r *ContactRepo) GetByID(conatctID int) (*amocrm.Contact, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return nil, nil
}

func (r *ContactRepo) GetAll() ([]amocrm.Contact, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return nil, nil
}

func (r *ContactRepo) Update(contact *amocrm.Contact) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	return nil
}

func (r *ContactRepo) Delete(contactID int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	return nil
}
