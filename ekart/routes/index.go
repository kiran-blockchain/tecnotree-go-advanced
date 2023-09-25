package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kiran-blockchain/ekart/controllers"
	"github.com/kiran-blockchain/ekart/middleware"
)


func UserRoutes(r *gin.Engine,a controllers.AuthController){
	//user routes
	user:= r.Group("/api/user")
	user.POST("/register",a.Register)
	user.POST("/login",a.Login)
}
func AuthRoutes(incomingRoutes *gin.Engine, a controllers.AuthController) {
    incomingRoutes.Use(middleware.Authenticate())
    incomingRoutes.GET("/usersdata", a.GetUser())
}
func ProductRoutes(r *gin.Engine,p controllers.ProductController){
	product:= r.Group("/api/product") //localhost:4000/api/product/
	product.POST("/insert",p.InsertProduct)
	product.GET("/getProducts",p.GetProducts)
	product.POST("/update",p.UpdateProduct)
}