package inmemorydb

import (
	"errors"
	"sync"

	"amocrm2.0/internal/core/amocrm"
)

var (
	errAccountNotExists = errors.New("such accountID does not exist")
)

type AccountRepo struct {
	accounts map[int]amocrm.Account
	mu       sync.RWMutex
}

func NewAccountRepo() *AccountRepo {
	return &AccountRepo{
		accounts: make(map[int]amocrm.Account),
	}
}

func (r *AccountRepo) Add(account *amocrm.Account) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.accounts[account.AccountID] = *account
	return nil
}

func (r *AccountRepo) GetByID(accountID int) (*amocrm.Account, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	val, ok := r.accounts[accountID]
	if !ok {
		return nil, errAccountNotExists
	}
	return &val, nil
}

func (r *AccountRepo) GetAll() ([]amocrm.Account, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	accounts := make([]amocrm.Account, 0, len(r.accounts))

	for _, val := range r.accounts {
		accounts = append(accounts, val)
	}
	return accounts, nil
}

func (r *AccountRepo) Update(account *amocrm.Account) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.accounts[account.AccountID]
	if !ok {
		return errAccountNotExists
	}

	r.accounts[account.AccountID] = *account
	return nil
}

func (r *AccountRepo) Delete(accountID int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.accounts[accountID]
	if !ok {
		return errAccountNotExists
	}

	delete(r.accounts, accountID)
	return nil
}
