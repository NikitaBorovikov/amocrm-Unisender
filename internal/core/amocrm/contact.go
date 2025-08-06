package amocrm

type Contact struct {
	ContactID  int    `gorm:"primaryKey"`
	AccountID  int    `gorm:"type:bigint"`
	Name       string `gorm:"type:varchar(255)"`
	Email      string `gorm:"type:varchar(255)"`
	SyncStatus bool   `gorm:"type:tinyint(1)"`
}
