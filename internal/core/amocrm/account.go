package amocrm

import "time"

type Account struct {
	AccountID    int
	Domain       string
	AccessToken  string
	RefreshToken string
	Expires      int64
	IssuedAt     time.Time
	UnisenderKey string
	Integration  []Integration
	Contacts     []Contact
}
