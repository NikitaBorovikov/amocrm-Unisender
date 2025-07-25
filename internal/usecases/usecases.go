package usecases

import "amocrm2.0/internal/core/amocrm"

type UseCases struct {
	AccountUC     *AccountUC
	IntegrationUC *IntegrationUC
	ContactUC     *ContactUC
}

func NewUseCases(accountRepo amocrm.AccountRepo, integrationRepo amocrm.IntegrationRepo, contactRepo amocrm.ContactRepo) *UseCases {
	return &UseCases{
		AccountUC:     NewAccountUC(accountRepo),
		IntegrationUC: NewIntegrationUC(integrationRepo),
		ContactUC:     NewContactUC(contactRepo),
	}
}
