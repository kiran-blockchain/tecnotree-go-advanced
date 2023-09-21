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
	
	
	
	
	//Cart routes
	//shipping routes
	// reporting route
}

func ProductRoutes(r *gin.Engine,p controllers.ProductController){
	product:= r.Group("/api/product") //localhost:4000/api/product/
	product.POST("/insert",p.InsertProduct)
	product.GET("/getProducts",p.GetProducts)
}