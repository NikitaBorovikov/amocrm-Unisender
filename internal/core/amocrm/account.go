package amocrm

type Account struct {
	AccountID    int
	AccessToken  string
	RefreshToken string
	Expires      int64
	UnisenderKey string
	Integration  []Integration
	Contacts     []Contact
}
