package mysqldb

import (
	"sync"

	"amocrm2.0/internal/core/amocrm"
	"gorm.io/gorm"
)

type AccountRepo struct {
	db *gorm.DB
	mu sync.RWMutex
}

func NewAccountRepo(db *gorm.DB) *AccountRepo {
	return &AccountRepo{db: db}
}

func (r *AccountRepo) Add(account *amocrm.Account) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	return nil
}

func (r *AccountRepo) GetByID(accountID int) (*amocrm.Account, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return nil, nil
}

func (r *AccountRepo) GetAll() ([]amocrm.Account, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return nil, nil
}

func (r *AccountRepo) Update(account *amocrm.Account) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	return nil
}

func (r *AccountRepo) Delete(accountID int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	return nil
}
