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
	return uc.AccountRepo.Add(account)
}

func (uc *AccountUC) GetByID(accountID int) (*amocrm.Account, error) {
	account, err := uc.AccountRepo.GetByID(accountID)
	if err != nil || account == nil {
		return nil, err
	}
	return account, nil
}

func (uc *AccountUC) GetAll() ([]amocrm.Account, error) {
	accounts, err := uc.AccountRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (uc *AccountUC) Update(account *amocrm.Account) error {
	return uc.AccountRepo.Update(account)
}

func (uc *AccountUC) Delete(accountID int) error {
	return uc.AccountRepo.Delete(accountID)
}
