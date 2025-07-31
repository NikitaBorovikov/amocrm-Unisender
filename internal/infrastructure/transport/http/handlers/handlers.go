package handlers

import (
	"fmt"
	"net/http"

	"amocrm2.0/internal/config"
	"amocrm2.0/internal/infrastructure/transport/http/dto"
	"amocrm2.0/internal/usecases"
	"github.com/go-chi/render"
)

type Handlers struct {
	AccountHandlers *AccountHandlers
	ContactHandlers *ContactHandlers
	Cfg             *config.Config
}

func NewHandlers(uc *usecases.UseCases, cfg *config.Config) *Handlers {
	return &Handlers{
		AccountHandlers: newAccountHandlers(uc.AccountUC, cfg),
		ContactHandlers: newContactHandlers(uc.ContactUC, uc.AccountUC),
		Cfg:             cfg,
	}
}

// functions for sending responses
func sendOKResponse(w http.ResponseWriter, r *http.Request, statusCode int, data interface{}, msg string) {
	w.WriteHeader(statusCode)
	render.JSON(w, r, dto.NewOKReponse(data, msg))
}

func sendErrorResponse(w http.ResponseWriter, r *http.Request, statusCode int, err error) {
	w.WriteHeader(statusCode)
	render.JSON(w, r, dto.NewErrorResponse(err))
}

// function for sending GET requests to amoCRM
func sendGetRequestToAmoCRM(accessToken, url string) (*http.Response, error) {
	req, err := prepareAmoCRMGetRequest(accessToken, url)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare request: %v", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("amoCRM error: status code %d", resp.StatusCode)
	}
	return resp, nil
}

func prepareAmoCRMGetRequest(accessToken, url string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)
	req.Header.Add("Accept", "application/json")
	return req, nil
}
