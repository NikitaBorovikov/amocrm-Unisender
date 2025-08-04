package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"amocrm2.0/internal/config"
	"amocrm2.0/internal/core/amocrm"
	"amocrm2.0/internal/infrastructure/queue"
	mysqldb "amocrm2.0/internal/infrastructure/repository/mysqlDB"
	"amocrm2.0/internal/infrastructure/transport/http/handlers"
	"amocrm2.0/internal/usecases"
	"amocrm2.0/internal/worker"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	contactTube = "contact_sync"
)

func main() {
	var workerID int
	flag.IntVar(&workerID, "worker-id", 0, "Worker ID")
	flag.Parse()

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

	err = db.AutoMigrate(
		&amocrm.Account{},
		&amocrm.Contact{},
	)
	if err != nil {
		logrus.Fatalf("failed to run DB migrate: %v", err)
	}

	producer := queue.NewProducer(beanstalk.Conn, contactTube)
	repo := mysqldb.NewMysqlRepo(db)
	usecases := usecases.NewUseCases(repo.AccountRepo, repo.ContactRepo)
	handlers := handlers.NewHandlers(usecases, producer, cfg)
	worker := worker.NewWorker(workerID, handlers, producer)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigChan
		logrus.Infof("Received signal: %v", sig)
		cancel()
	}()

	worker.Run(ctx)
}

func initMySQL(cfg *config.DB) (*gorm.DB, error) {
	dsn := makeDSN(cfg)
	db, err := gorm.Open(mysql.Open(dsn))
	return db, err
}

func makeDSN(cfg *config.DB) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
}
