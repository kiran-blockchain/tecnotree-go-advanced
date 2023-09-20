package controllers

import (
	"github.com/gin-gonic/gin"
)


func HandleLogin(c *gin.Context) {
	
	c.JSON(200, gin.H{
		"data": "Welcome to login",
	})
}

func HandleRegister(c *gin.Context) {
	
	c.JSON(200, gin.H{
		"data": "Register",
	})
}
func HandleGetProfile(c *gin.Context) {
	
	c.JSON(200, gin.H{
		"data": "Profile",
	})
}

func HandleLogout(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Logged out",
	})
}

// fetch user details
//1. attach the bearer token as part of header
//2. Intercept the token and decrypt it -- middleware
//3. Extract the userid from the decrypted token
//4. query the data base with this userid
//5. return the details
