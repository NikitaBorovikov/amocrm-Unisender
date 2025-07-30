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

	result := r.db.Create(integration)
	return result.Error
}

func (r *IntegrationRepo) GetByID(integrationID int) (*amocrm.Integration, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var integration amocrm.Integration
	result := r.db.Where("integration_id = ?", integrationID).First(&integration)
	return &integration, result.Error
}

func (r *IntegrationRepo) GetAll() ([]amocrm.Integration, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var integrations []amocrm.Integration
	result := r.db.Find(&integrations)
	return integrations, result.Error
}

func (r *IntegrationRepo) Update(integration *amocrm.Integration) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	result := r.db.Model(integration).Updates(integration)
	return result.Error
}

func (r *IntegrationRepo) Delete(integrationID int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	result := r.db.Delete(&amocrm.Integration{}, integrationID)
	return result.Error
}
