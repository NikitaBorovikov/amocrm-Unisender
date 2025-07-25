package handlers

import "amocrm2.0/internal/usecases"

type Handlers struct {
	AccountHandlers     *AccountHandlers
	IntegrationHandlers *IntegrationHandlers
	ContactHandlers     *ContactHandlers
}

func NewHandlers(uc *usecases.UseCases) *Handlers {
	return &Handlers{
		AccountHandlers:     newAccountHandlers(uc.AccountUC),
		IntegrationHandlers: newIntegrationHandlers(uc.IntegrationUC),
		ContactHandlers:     newContactHandlers(uc.ContactUC),
	}
}
