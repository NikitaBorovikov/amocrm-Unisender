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

func (r *ContactRepo) Update(contact *amocrm.Contact) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	result := r.db.Model(contact).Updates(contact)
	return result.Error
}
