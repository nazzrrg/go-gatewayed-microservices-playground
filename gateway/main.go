package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nazzrrg/go-gatewayed-microservice-playgroud/gateway/controllers"
	"github.com/nazzrrg/go-gatewayed-microservice-playgroud/gateway/models"
	"log"
	"os"
)

func main() {
	models.ConnectToDb()
	router := gin.Default()

	api := router.Group("/api/v1")
	auth := api.Group("/auth")
	auth.POST("/register", controllers.Register)
	auth.POST("/login", controllers.Login)

	//todo: gRPC calls to other services with forwarding
	//users := api.Group("/users")

	address, ok := os.LookupEnv("GATEWAY_PORT")
	if !ok {
		address = "3000"
	}
	address = ":" + address
	if err := router.Run(address); err != nil {
		log.Fatalf("Error running server on address %s\n error: %v\n", address, err)
	}
}
