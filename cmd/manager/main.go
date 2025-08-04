package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:  "amo-sync",
		Usage: "Сервис синхронизации контактов amoCRM с Unisender",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:     "workers",
				Usage:    "Number of workers to start",
				Required: true,
				Value:    1,
			},
		},
		Action: func(c *cli.Context) error {
			numWorkers := c.Int("workers")
			workerPath := "./workers"
			if _, err := os.Stat(workerPath); err != nil {
				return fmt.Errorf("failed to find worker binary file: %v", err)
			}

			workersDone := make(chan error, numWorkers)
			var wg sync.WaitGroup

			for i := 0; i < numWorkers; i++ {
				wg.Add(1)
				go func(workerID int) {
					defer wg.Done()
					cmd := exec.Command(workerPath, "--worker-id", fmt.Sprintf("%d", i))

					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr

					if err := cmd.Start(); err != nil {
						logrus.Errorf("failed to start worker %d: %v", i, err)
						return
					}

					logrus.Infof("Started worker %d (PID: %d)", i, cmd.Process.Pid)

					if err := cmd.Wait(); err != nil {
						workersDone <- fmt.Errorf("worker %d exited: %w", workerID, err)
					}
				}(i)
			}
			go func() {
				sigChan := make(chan os.Signal, 1)
				signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
				<-sigChan
				logrus.Info("Received shutdown signal")
				close(workersDone)
			}()

			wg.Wait()
			close(workersDone)

			for err := range workersDone {
				if err != nil {
					logrus.Error(err)
				}
			}
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
