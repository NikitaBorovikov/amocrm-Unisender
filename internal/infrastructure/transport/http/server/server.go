package server

import (
	"net/http"

	"amocrm2.0/internal/infrastructure/transport/http/handlers"
	"github.com/sirupsen/logrus"
)

func Run(h *handlers.Handlers, port string) {
	initRoutes(h)

	logrus.Infof("http-server is starting on port %s...", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		logrus.Fatal(err)
	}
}
