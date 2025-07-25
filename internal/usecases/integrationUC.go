package usecases

import "amocrm2.0/internal/core/amocrm"

type IntegrationUC struct {
	IntegrationRepo amocrm.IntegrationRepo
}

func NewIntegrationUC(integrationRepo amocrm.IntegrationRepo) *IntegrationUC {
	return &IntegrationUC{
		IntegrationRepo: integrationRepo,
	}
}

func (uc *IntegrationUC) Add(integration *amocrm.Integration) error {
	return nil
}

func (uc *IntegrationUC) GetByID(integrationID int) (*amocrm.Integration, error) {
	return nil, nil
}

func (uc *IntegrationUC) GetAll() ([]amocrm.Integration, error) {
	return nil, nil
}

func (uc *IntegrationUC) Update(integration *amocrm.Integration) error {
	return nil
}

func (uc *IntegrationUC) Delete(integrationID int) error {
	return nil
}
