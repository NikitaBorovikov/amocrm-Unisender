package amocrm

import "time"

type Account struct {
	AccountID    int `gorm:"primaryKey"`
	Domain       string
	AccessToken  string
	RefreshToken string
	Expires      int64
	IssuedAt     time.Time
	UnisenderKey string
	Contacts     []Contact `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE"`
}
