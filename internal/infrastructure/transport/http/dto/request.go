package dto

type IntegrationInfoRequest struct {
	AuthCode string
	Domain   string
}

type ExchangeTokensRequest struct {
	ClientID    string `json:"client_id"`
	SecretKey   string `json:"client_secret"`
	GrantType   string `json:"grant_type"`
	Code        string `json:"code"`
	RedirectURL string `json:"redirect_uri"`
}

type RefreshAccessTokenRequest struct {
	ClinetID     string `json:"client_id"`
	SecretKey    string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
	RedirectURL  string `json:"redirect_uri"`
}

type UnisenderImportRequest struct {
	APIKey     string     `json:"api_key"`
	FieldNames []string   `json:"field_names"`
	Data       [][]string `json:"data"`
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

func NewExchangeTokensRequest(code, clientID, secretKey, redirectURL, grantType string) *ExchangeTokensRequest {
	return &ExchangeTokensRequest{
		Code:        code,
		ClientID:    clientID,
		SecretKey:   secretKey,
		RedirectURL: redirectURL,
		GrantType:   grantType,
	}
}

func NewRefreshAccessTokenRequest(clinetID, secretKey, redirectURL, refreshToken, grantType string) *RefreshAccessTokenRequest {
	return &RefreshAccessTokenRequest{
		ClinetID:     clinetID,
		SecretKey:    secretKey,
		RefreshToken: refreshToken,
		RedirectURL:  redirectURL,
		GrantType:    grantType,
	}
}

func NewIntegrationInfo(authCode, domain string) *IntegrationInfoRequest {
	return &IntegrationInfoRequest{
		AuthCode: authCode,
		Domain:   domain,
	}
}
