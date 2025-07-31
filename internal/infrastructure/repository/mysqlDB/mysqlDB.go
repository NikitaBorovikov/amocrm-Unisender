package mysqldb

import "gorm.io/gorm"

type MySqlRepo struct {
	AccountRepo *AccountRepo
	ContactRepo *ContactRepo
}

func NewMysqlRepo(db *gorm.DB) *MySqlRepo {
	return &MySqlRepo{
		AccountRepo: NewAccountRepo(db),
		ContactRepo: NewContactRepo(db),
	}
}
