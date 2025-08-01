package usecases

import "amocrm2.0/internal/core/amocrm"

type UseCases struct {
	AccountUC *AccountUC
	ContactUC *ContactUC
}

func NewUseCases(accountRepo amocrm.AccountRepo, contactRepo amocrm.ContactRepo) *UseCases {
	return &UseCases{
		AccountUC: NewAccountUC(accountRepo),
		ContactUC: NewContactUC(contactRepo),
	}
}
