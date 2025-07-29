package server

import (
	"net/http"

	"amocrm2.0/internal/infrastructure/transport/http/handlers"
)

func initRoutes(h *handlers.Handlers) {
	//r := chi.NewRouter()

	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		h.HandleAuth(w, r)
	})

	http.HandleFunc("/contacts", func(w http.ResponseWriter, r *http.Request) {
		h.GetContacts(0) //REMOVE: FOR TESTING
	})
}
