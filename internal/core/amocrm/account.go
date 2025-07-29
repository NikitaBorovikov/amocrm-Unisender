package amocrm

import "time"

type Account struct {
	AccountID    int `gorm:"primartKey"`
	Domain       string
	AccessToken  string
	RefreshToken string
	Expires      int64
	IssuedAt     time.Time
	UnisenderKey string
	Integration  []Integration `gorm:"foreignKey:AccountID"`
	Contacts     []Contact     `gorm:"foreignKey:AccountID"`
}
