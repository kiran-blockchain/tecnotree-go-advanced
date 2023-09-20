package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kiran-blockchain/ekart/controllers"
)


func AppRoutes(r *gin.Engine){
	//user routes
	user:= r.Group("/api/user")

	user.POST("/register",controllers.HandleRegister)
	user.POST("/login",controllers.HandleLogin)
	user.GET("/logout",controllers.HandleLogout)
	user.GET("/profile/:id",controllers.HandleGetProfile)
	//product routes
	product:= r.Group("/api/product") //localhost:4000/api/product/
	
	product.POST("/add",controllers.HandleRegister)
	product.POST("/edit",controllers.HandleLogin)
	product.GET("/search",controllers.HandleLogout)
	product.GET("/product/:id",controllers.HandleGetProfile)
	//Cart routes
	//shipping routes
	// reporting route
}