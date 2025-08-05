package grpc

import (
	"net"

	"amocrm2.0/internal/usecases"
	pb "amocrm2.0/proto"
	"google.golang.org/grpc"
)

func RunGRPCServer(port string, uc *usecases.AccountUC) error {
	server := grpc.NewServer()
	pb.RegisterAccountServiceServer(server, uc)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	return server.Serve(lis)
}
