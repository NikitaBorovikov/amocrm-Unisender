package app

import (
	"context"
	"time"

	"amocrm2.0/internal/config"
	pb "amocrm2.0/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RunClient() {
	cfg, err := config.InitConfig()
	if err != nil {
		logrus.Fatalf("failed to init config: %v", err)
	}

	serverAddr := cfg.GrpcServer.Host + cfg.GrpcServer.Port
	conn, err := grpc.NewClient(
		serverAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logrus.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewAccountServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	response, err := client.Unsubscribe(ctx, &pb.UnsubscribeRequest{AccountId: 32573390})
	if err != nil {
		logrus.Fatalf("failed to unsubscribe: %v", err)
		return
	}
	logrus.Infof("successful unsubscribe: status = %v, message = %s", response.Success, response.Message)
}
