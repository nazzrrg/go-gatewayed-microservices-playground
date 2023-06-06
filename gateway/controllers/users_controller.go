package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/nazzrrg/go-gatewayed-microservice-playgroud/gateway/user_service"
	"log"
	"net/http"
)

type UserController struct {
	GrpcClient user_service.UserServiceClient
}

type deleteUserInput struct {
	Id int `json:"id" binding:"required"`
}

func (controller *UserController) DeleteUser(c *gin.Context) {
	var input deleteUserInput

	var err error
	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body format"})
		return
	}

	response, err := controller.GrpcClient.DeleteUser(context.Background(), &user_service.DeleteUserRequest{Id: uint32(input.Id)})
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to delete user"})
		return
	}
	if !response.Success {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to delete user from user service"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
