package server

import (
	"finnbank/services/account/service"
	"finnbank/services/common/grpc/account"
	"finnbank/services/common/utils"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

// Set up grpc connection

func GrpcServer(s service.AccountService, logger *utils.Logger) error {
	lis, err := net.Listen("tcp", "localhost:8082")
	if err != nil {
		logger.Fatal("Could not start gRPC server on port 9000: %s", err)
		return err
	}
	logger.Info("Port 8082 listening success")
	grpcServer := grpc.NewServer()
	account.RegisterAccountServiceServer(grpcServer, &s)

	// Graceful Shutdown Handling
	// Definitely didnt copy this from stackoverflow :D
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		logger.Info("Shutting down gRPC server...")
		grpcServer.GracefulStop()
		lis.Close()
		os.Exit(0)
	}()
	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatal("Failed to start gRPC server: %s", err)
		return err
	}
	return nil
}
