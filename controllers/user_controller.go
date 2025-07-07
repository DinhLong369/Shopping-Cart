package controllers

import (
	"Shopping-cart/models"
	"Shopping-cart/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service}
}

func (ctrl *UserController) SignUp(c *gin.Context) {
	var input models.User
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err := ctrl.service.SignUp(&input)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
	}
	if err == nil {
		c.JSON(http.StatusCreated, gin.H{"message": "Dang ky thanh cong"})
	}
}

func (ctrl *UserController) Login(c *gin.Context) {
	var input models.LoginUser
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	token, err := ctrl.service.Login(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
	}
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Dang nhap thanh cong", "token": token})
	}
}
