package app

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"amocrm2.0/internal/config"
	"amocrm2.0/internal/core/amocrm"
	"amocrm2.0/internal/infrastructure/queue"
	mysqldb "amocrm2.0/internal/infrastructure/repository/mysqlDB"
	"amocrm2.0/internal/infrastructure/transport/http/handlers"
	"amocrm2.0/internal/infrastructure/transport/http/server"
	"amocrm2.0/internal/usecases"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	contactTube = "contact_sync"
)

func RunServer() {
	cfg, err := config.InitConfig()
	if err != nil {
		logrus.Fatalf("failed to init config: %v", err)
	}

	beanstalk, err := queue.InitBeanstalk(&cfg.Beanstalk)
	if err != nil {
		logrus.Fatalf("failed to init beanstalk: %v", err)
	}
	defer beanstalk.Conn.Close()

	db, err := initMySQL(&cfg.DB)
	if err != nil {
		logrus.Fatalf("failed to init DB: %v", err)
	}

	if err := db.AutoMigrate(&amocrm.Account{}, &amocrm.Contact{}); err != nil {
		logrus.Fatalf("failed to run DB migrate: %v", err)
	}

	//repo := inmemorydb.NewInmomryDB()
	producer := queue.NewProducer(beanstalk.Conn, contactTube)
	repo := mysqldb.NewMysqlRepo(db)
	usecases := usecases.NewUseCases(repo.AccountRepo, repo.ContactRepo)
	handlers := handlers.NewHandlers(usecases, producer, cfg)

	var wg sync.WaitGroup
	serverErrors := make(chan error, 2)

	// // //grpc-server
	// // wg.Add(1)
	// // go func() {
	// // 	defer wg.Done()
	// // }()

	wg.Add(1)
	go func() {
		defer wg.Done()
		logrus.Infof("http-server is starting on port %s...", cfg.RestServer.Port)
		if err := server.Run(handlers, cfg.RestServer.Port); err != nil {
			serverErrors <- err
		}

	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-quit:
		logrus.Info("Received shutdown signal")
	case err := <-serverErrors:
		logrus.Errorf("Server error: %v", err)
	}

	wg.Wait()
	// go server.Run(handlers, cfg.RestServer.Port)

	// select {}
}

func initMySQL(cfg *config.DB) (*gorm.DB, error) {
	dsn := makeDSN(cfg)
	db, err := gorm.Open(mysql.Open(dsn))
	return db, err
}

func makeDSN(cfg *config.DB) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
}
