package app

import (
	"amocrm2.0/internal/config"
	inmemorydb "amocrm2.0/internal/infrastructure/repository/inmemoryDB"
	"amocrm2.0/internal/infrastructure/transport/http/handlers"
	"amocrm2.0/internal/infrastructure/transport/http/server"
	"amocrm2.0/internal/usecases"
	"github.com/sirupsen/logrus"
)

func RunServer() {
	cfg, err := config.InitConfig()
	if err != nil {
		logrus.Fatalf("failed to init config: %v", err)
	}

	repo := inmemorydb.NewInmomryDB()
	usecase := usecases.NewUseCases(repo.AccountRepo, repo.IntegrationRepo, repo.ContactRepo)
	handlers := handlers.NewHandlers(usecase, cfg)
	go server.Run(handlers, cfg.RestServer.Port)

	select {}
}
