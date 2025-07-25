package inmemorydb

import (
	"errors"
	"sync"

	"amocrm2.0/internal/core/amocrm"
)

var (
	errIntegrationNotExists = errors.New("such integrationID does not exist")
)

type IntegrationRepo struct {
	integrations map[int]amocrm.Integration
	mu           sync.RWMutex
}

func NewIntegrationRepo() *IntegrationRepo {
	return &IntegrationRepo{
		integrations: make(map[int]amocrm.Integration),
	}
}

func (r *IntegrationRepo) Add(integration *amocrm.Integration) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.integrations[integration.IntegrationID] = *integration
	return nil
}

func (r *IntegrationRepo) GetByID(integrationID int) (*amocrm.Integration, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	val, ok := r.integrations[integrationID]
	if !ok {
		return nil, errIntegrationNotExists
	}
	return &val, nil
}

func (r *IntegrationRepo) GetAll() ([]amocrm.Integration, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	intergations := make([]amocrm.Integration, 0, len(r.integrations))

	for _, val := range r.integrations {
		intergations = append(intergations, val)
	}
	return intergations, nil
}

func (r *IntegrationRepo) Update(integration *amocrm.Integration) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.integrations[integration.IntegrationID]
	if !ok {
		return errIntegrationNotExists
	}
	r.integrations[integration.IntegrationID] = *integration
	return nil
}

func (r *IntegrationRepo) Delete(integrationID int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.integrations[integrationID]
	if !ok {
		return errIntegrationNotExists
	}

	delete(r.integrations, integrationID)
	return nil
}
