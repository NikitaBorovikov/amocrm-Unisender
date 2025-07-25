package app

import (
	inmemorydb "amocrm2.0/internal/infrastructure/repository/inmemoryDB"
	"amocrm2.0/internal/usecases"
)

func RunServer() {
	//TODO: init config

	repo := inmemorydb.NewInmomryDB()
	usecase := usecases.NewUseCases(repo.AccountRepo, repo.IntegrationRepo, repo.ContactRepo)
	_ = usecase //REMOVE
}
