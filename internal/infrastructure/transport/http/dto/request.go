package dto

type AddAccountRequest struct {
	AccountID    int    `json:"account_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Expires      int64  `json:"expires"`
}

type UpdateAccountRequest struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Expires      int64  `json:"expires"`
}

type AddIntegrationRequest struct {
	IntegraionID int    `json:"integration_id"`
	AccountID    int    `json:"account_id"`
	ClinetID     string `json:"client_id"`
	SecretKey    string `json:"secret_key"`
	RedirectURL  string `json:"redirect_url"`
	AuthCode     string `json:"auth_code"`
}

type UpdateIntegrationRequest struct {
	AccountID   int    `json:"account_id"`
	ClinetID    string `json:"client_id"`
	SecretKey   string `json:"secret_key"`
	RedirectURL string `json:"redirect_url"`
	AuthCode    string `json:"auth_code"`
}
