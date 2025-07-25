package amocrm

type Integration struct {
	IntegrationID int
	AccountID     int
	AuthCode      string
	ClientID      string
	SecretKey     string
	RedirectURL   string
}
