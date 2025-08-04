package server

import (
	"net/http"

	"amocrm2.0/internal/infrastructure/transport/http/handlers"
)

func Run(h *handlers.Handlers, port string) error {
	initRoutes(h)
	return http.ListenAndServe(port, nil)
}
