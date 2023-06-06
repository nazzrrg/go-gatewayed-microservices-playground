package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nazzrrg/go-gatewayed-microservice-playgroud/gateway/controllers"
	"github.com/nazzrrg/go-gatewayed-microservice-playgroud/gateway/models"
	"github.com/nazzrrg/go-gatewayed-microservice-playgroud/gateway/user_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

func main() {
	models.ConnectToDb()
	router := gin.Default()

	usHost, ok := os.LookupEnv("USER_SERVICE_HOST")
	if !ok {
		usHost = "localhost"
	}
	usHost = usHost + ":"
	usPort, ok := os.LookupEnv("USER_SERVICE_PORT")
	if ok {
		usHost = usHost + usPort
	} else {
		usHost = usHost + "9000"
	}

	conn, err := grpc.Dial(usHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial grpc server")
	}
	usersGrpcClient := user_service.NewUserServiceClient(conn)

	api := router.Group("/api/v1")
	auth := api.Group("/auth")
	auth.POST("/register", controllers.Register)
	auth.POST("/login", controllers.Login)

	userController := controllers.UserController{GrpcClient: usersGrpcClient}
	//users := api.Group("/users")
	api.DELETE("/users", userController.DeleteUser)
	//users.POST("/", userController.CreateUser)
	//users.PUT("/", userController.UpdateUser)

	address, ok := os.LookupEnv("GATEWAY_PORT")
	if !ok {
		address = "3000"
	}
	address = ":" + address
	if err := router.Run(address); err != nil {
		log.Fatalf("Error running server on address %s\n error: %v\n", address, err)
	}
}
