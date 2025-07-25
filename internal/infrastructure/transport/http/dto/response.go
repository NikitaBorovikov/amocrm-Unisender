package dto

type ErrorResponse struct {
	Msg string `json:"error"`
}

type OKResponse struct {
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
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
