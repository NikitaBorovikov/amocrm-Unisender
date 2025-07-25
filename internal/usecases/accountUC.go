package usecases

import "amocrm2.0/internal/core/amocrm"

type AccountUC struct {
	AccountRepo amocrm.AccountRepo
}

func NewAccountUC(accountRepo amocrm.AccountRepo) *AccountUC {
	return &AccountUC{
		AccountRepo: accountRepo,
	}
}

func (uc *AccountUC) Add(account *amocrm.Account) error {
	return nil
}

func (uc *AccountUC) GetByID(accountID int) (*amocrm.Account, error) {
	return nil, nil
}

func (uc *AccountUC) GetAll() ([]amocrm.Account, error) {
	return nil, nil
}

func (uc *AccountUC) Update(account *amocrm.Account) error {
	return nil
}

func (uc *AccountUC) Delete(accountID int) error {
	return nil
}
