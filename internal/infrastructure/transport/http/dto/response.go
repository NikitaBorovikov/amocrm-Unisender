package dto

import "amocrm2.0/internal/core/amocrm"

type ErrorResponse struct {
	Msg string `json:"error"`
}

type OKResponse struct {
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

type ExchangeTokensResponse struct {
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Expires      int64  `json:"expires_in"`
}

type RefreshAccessTokenResponse struct {
	TokenType    string `json:"token_type"`
	Expires      int64  `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{
		Msg: err.Error(),
	}
}

func NewOKReponse(data interface{}, msg string) *OKResponse {
	return &OKResponse{
		Msg:  msg,
		Data: data,
	}
}

func (r *ExchangeTokensResponse) ToDomainAccount() amocrm.Account {
	return amocrm.Account{
		AccessToken:  r.AccessToken,
		RefreshToken: r.RefreshToken,
		Expires:      r.Expires,
	}
}
