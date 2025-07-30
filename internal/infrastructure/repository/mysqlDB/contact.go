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

	result := r.db.Create(contact)
	return result.Error
}

func (r *ContactRepo) GetByID(contactID int) (*amocrm.Contact, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var contact amocrm.Contact
	result := r.db.Where("contact_id = ?", contactID).First(&contact)
	return &contact, result.Error
}

func (r *ContactRepo) GetAll() ([]amocrm.Contact, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var contacts []amocrm.Contact
	result := r.db.Find(&contacts)
	return contacts, result.Error
}

func (r *ContactRepo) Update(contact *amocrm.Contact) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	result := r.db.Model(contact).Updates(contact)
	return result.Error
}

func (r *ContactRepo) Delete(contactID int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	result := r.db.Delete(&amocrm.Contact{}, contactID)
	return result.Error
}
