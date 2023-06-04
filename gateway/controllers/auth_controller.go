package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nazzrrg/go-gatewayed-microservice-playgroud/gateway/models"
	"github.com/nazzrrg/go-gatewayed-microservice-playgroud/gateway/utils"
	"net/http"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput

	var err error
	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body format"})
		return
	}

	u, err := models.FindByUsername(input.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user does not exist"})
		return
	}

	if err = utils.VerifyPassword(u.Password, input.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		return
	}

	token, err := utils.GenerateToken(u.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": token})
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body format"})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	_, err := u.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to create user. User may already exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
