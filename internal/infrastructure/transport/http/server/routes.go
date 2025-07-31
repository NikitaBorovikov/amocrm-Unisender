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

	// REMOVE: FOR TESTING
	http.HandleFunc("/sync", func(w http.ResponseWriter, r *http.Request) {
		h.HandleFirstSync(32573390) //REMOVE: FOR TESTING
	})

	http.HandleFunc("/api_key", func(w http.ResponseWriter, r *http.Request) {
		h.ReceiveUnisenderKey(w, r)
	})
}
