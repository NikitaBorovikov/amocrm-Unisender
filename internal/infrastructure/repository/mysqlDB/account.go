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

	result := r.db.Create(account)
	return result.Error
}

func (r *AccountRepo) GetByID(accountID int) (*amocrm.Account, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var account amocrm.Account
	result := r.db.Where("account_id = ?", accountID).First(&account)
	return &account, result.Error
}

func (r *AccountRepo) GetAll() ([]amocrm.Account, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var accounts []amocrm.Account
	result := r.db.Find(&accounts)
	return accounts, result.Error
}

func (r *AccountRepo) Update(account *amocrm.Account) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	result := r.db.Model(account).Updates(account)
	return result.Error
}

func (r *AccountRepo) Delete(accountID int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	result := r.db.Delete(&amocrm.Account{}, accountID)
	return result.Error
}
