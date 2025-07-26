package server

import (
	"net/http"

	"amocrm2.0/internal/infrastructure/transport/http/handlers"
)

func initRoutes(h *handlers.Handlers) {
	//r := chi.NewRouter()

	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		h.AccountHandlers.HandleAuth(w, r)
	})
}
