package mysqldb

import "gorm.io/gorm"

type MySqlRepo struct {
	AccountRepo     *AccountRepo
	IntegrationRepo *IntegrationRepo
	ContactRepo     *ContactRepo
}

func NewMysqlRepo(db *gorm.DB) *MySqlRepo {
	return &MySqlRepo{
		AccountRepo:     NewAccountRepo(db),
		IntegrationRepo: NewIntegrationRepo(db),
		ContactRepo:     NewContactRepo(db),
	}
}
