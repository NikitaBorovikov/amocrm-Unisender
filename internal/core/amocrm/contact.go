package amocrm

type Contact struct {
	ContactID  int `gorm:"primaryKey"`
	AccountID  int
	Name       string
	Email      string
	SyncStatus bool
}
