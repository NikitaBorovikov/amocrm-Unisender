package app

import (
	"fmt"

	"amocrm2.0/internal/config"
	"amocrm2.0/internal/core/amocrm"
	mysqldb "amocrm2.0/internal/infrastructure/repository/mysqlDB"
	"amocrm2.0/internal/infrastructure/transport/http/handlers"
	"amocrm2.0/internal/infrastructure/transport/http/server"
	"amocrm2.0/internal/usecases"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func RunServer() {
	cfg, err := config.InitConfig()
	if err != nil {
		logrus.Fatalf("failed to init config: %v", err)
	}

	db, err := initMySQL(&cfg.DB)
	if err != nil {
		logrus.Fatalf("failed to init DB: %v", err)
	}

	err = db.AutoMigrate(
		&amocrm.Account{},
		&amocrm.Contact{},
	)
	if err != nil {
		logrus.Fatalf("failed to run DB migrate: %v", err)
	}

	//repo := inmemorydb.NewInmomryDB()
	repo := mysqldb.NewMysqlRepo(db)
	usecases := usecases.NewUseCases(repo.AccountRepo, repo.ContactRepo)
	handlers := handlers.NewHandlers(usecases, cfg)
	go server.Run(handlers, cfg.RestServer.Port)

	select {}
}

func initMySQL(cfg *config.DB) (*gorm.DB, error) {
	dsn := makeDSN(cfg)
	db, err := gorm.Open(mysql.Open(dsn))
	return db, err
}

func makeDSN(cfg *config.DB) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
}
