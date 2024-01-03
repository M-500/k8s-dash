//@Author: wulinlin
//@Description:
//@File:  user_controller
//@Version: 1.0.0
//@Date: 2024/01/02 21:58

package controller

import (
	dao2 "gin-server/app/dao"
	"gin-server/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (controller *UserController) Register(c *gin.Context) {
	var requestBody struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dao := dao2.NewUserDao(c)
	userService := services.NewUserService(dao, c)

	if err := userService.RegisterUser(requestBody.Username, requestBody.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (controller *UserController) Login(c *gin.Context) {
	var requestBody struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dao := dao2.NewUserDao(c)
	userService := services.NewUserService(dao, c)
	user, err := userService.AuthenticateUser(requestBody.Username, requestBody.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// 这里可以生成 JWT 签名等，根据需要进行处理

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
}

func (controller *UserController) Demo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "玩个头",
	})
}
