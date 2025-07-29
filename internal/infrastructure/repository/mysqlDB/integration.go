package mysqldb

import (
	"sync"

	"amocrm2.0/internal/core/amocrm"
	"gorm.io/gorm"
)

type IntegrationRepo struct {
	db *gorm.DB
	mu sync.RWMutex
}

func NewIntegrationRepo(db *gorm.DB) *IntegrationRepo {
	return &IntegrationRepo{db: db}
}

func (r *IntegrationRepo) Add(integration *amocrm.Integration) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	return nil
}

func (r *IntegrationRepo) GetByID(integrationID int) (*amocrm.Integration, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return nil, nil
}

func (r *IntegrationRepo) GetAll() ([]amocrm.Integration, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return nil, nil
}

func (r *IntegrationRepo) Update(integration *amocrm.Integration) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	return nil
}

func (r *IntegrationRepo) Delete(integrationID int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	return nil
}
