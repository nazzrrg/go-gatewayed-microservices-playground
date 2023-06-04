package main

import (
	"github.com/nazzrrg/go-gatewayed-microservices-playgroud/user_service/service"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	address, ok := os.LookupEnv("USER_SERVICE_PORT")
	if !ok {
		address = "9000"
	}
	address = ":" + address
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen on address %s: %v\n", address, err)
	}

	userService := service.UserService{}
	grpcServer := grpc.NewServer()
	service.RegisterUserServiceServer(grpcServer, &userService)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server on address %s: %v", address, err)
	}
}
