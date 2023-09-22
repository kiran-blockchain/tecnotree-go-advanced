package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kiran-blockchain/ekart/entities"
	"github.com/kiran-blockchain/ekart/interfaces"
)
type AuthController struct {
	AuthService interfaces.IUser
}

func InitAuthController(authService interfaces.IUser) *AuthController {
	return &AuthController{AuthService: authService}
}

func(a *AuthController) Register(c *gin.Context){
	fmt.Println("Invoked controller")
	var user entities.User
	err := c.BindJSON(&user)
	if err != nil {
		return
	}
	result, err := a.AuthService.Register(&user)
	fmt.Println(result)
	if err != nil {
		return
	} else {
		c.IndentedJSON(http.StatusCreated, result)
	}
}