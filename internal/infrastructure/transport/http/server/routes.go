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

	// Important: getting the Unisender API key is the trigger for starting the initial contact synchronization.
	http.HandleFunc("/api_key", func(w http.ResponseWriter, r *http.Request) {
		h.ReceiveUnisenderKey(w, r)
	})

	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		h.ReceiveContactWebhook(w, r)
	})
}
