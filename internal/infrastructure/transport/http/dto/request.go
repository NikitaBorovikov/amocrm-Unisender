package dto

type ExchangeTokensRequest struct {
	ClientID    string `json:"client_id"`
	SecretKey   string `json:"client_secret"`
	GrantType   string `json:"grant_type"`
	Code        string `json:"code"`
	RedirectURL string `json:"redirect_uri"`
}

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

func NewExchangeTokensRequest(clientID, secretKey, grantType, code, redirectURL string) *ExchangeTokensRequest {
	return &ExchangeTokensRequest{
		ClientID:    clientID,
		SecretKey:   secretKey,
		GrantType:   grantType,
		Code:        code,
		RedirectURL: redirectURL,
	}
}
