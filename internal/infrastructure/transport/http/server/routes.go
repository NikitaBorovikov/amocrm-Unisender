package server

import (
	"net/http"

	"amocrm2.0/internal/infrastructure/transport/http/handlers"
)

func initRoutes(h *handlers.Handlers) {
	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		h.HandleAuth(w, r)
	})

	// Получение Unisender API ключа является триггером для первичной синхронизации контактов
	http.HandleFunc("/api_key", func(w http.ResponseWriter, r *http.Request) {
		h.ReceiveUnisenderKey(w, r)
	})

	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		h.ReceiveContactWebhook(w, r)
	})
}
