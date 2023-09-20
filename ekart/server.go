package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kiran-blockchain/ekart/routes"
)

func main() {
	server := gin.Default()
	
	routes.AppRoutes(server)
	server.GET("/", home)
	server.Run(":4000")
}

func home(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Welcome to gin",
	})
}
