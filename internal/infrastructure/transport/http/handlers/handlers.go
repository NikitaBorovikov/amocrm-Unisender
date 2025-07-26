package handlers

import (
	"net/http"

	"amocrm2.0/internal/config"
	"amocrm2.0/internal/infrastructure/transport/http/dto"
	"amocrm2.0/internal/usecases"
	"github.com/go-chi/render"
)

type Handlers struct {
	AccountHandlers     *AccountHandlers
	IntegrationHandlers *IntegrationHandlers
	ContactHandlers     *ContactHandlers
}

func NewHandlers(uc *usecases.UseCases, cfg *config.Config) *Handlers {
	return &Handlers{
		AccountHandlers:     newAccountHandlers(uc.AccountUC, cfg),
		IntegrationHandlers: newIntegrationHandlers(uc.IntegrationUC),
		ContactHandlers:     newContactHandlers(uc.ContactUC),
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
