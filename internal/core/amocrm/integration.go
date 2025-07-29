package amocrm

type Integration struct {
	IntegrationID int `gorm:"primaryKey"`
	AccountID     int
	AuthCode      string
	ClientID      string
	SecretKey     string
	RedirectURL   string
}
