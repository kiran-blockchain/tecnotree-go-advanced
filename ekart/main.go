package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kiran-blockchain/ekart/config"
	"github.com/kiran-blockchain/ekart/controllers"
	"github.com/kiran-blockchain/ekart/routes"
	"github.com/kiran-blockchain/ekart/services"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoClient *mongo.Client
	err         error
	ctx         context.Context
	server      *gin.Engine
)

func main() {
	server = gin.Default()
	InitializeDatabase()
	InitializeAuthentication()
	InitializeProducts()
	ctx1, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer mongoClient.Disconnect(ctx1)
	server.Run(":4000")
}

func InitializeDatabase() {
	mongoClient, err = config.ConnectDataBase()
	if err != nil {
		log.Fatalf("Unable to connect to Database", err)
	} else {
		fmt.Println("Connected to Database")
	}
}
func InitializeProducts() {
	productCollection := config.GetCollection(mongoClient, "ekart", "products")
	productSvc := services.InitProductService(productCollection)
	productCtrl := controllers.InitProductController(productSvc)
	routes.ProductRoutes(server, *productCtrl)
}

func InitializeAuthentication() {
	collection := config.GetCollection(mongoClient, "ekart", "users")
	authSvc := services.InitUserService(collection)
	authCtrl := controllers.InitAuthController(authSvc)
	routes.UserRoutes(server, *authCtrl)
	routes.AuthRoutes(server, *authCtrl)

}
