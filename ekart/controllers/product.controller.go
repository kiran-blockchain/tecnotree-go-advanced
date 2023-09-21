package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kiran-blockchain/ekart/entities"
	"github.com/kiran-blockchain/ekart/interfaces"
)

type ProductController struct {
	ProductService interfaces.IProduct
}

func InitProductController(productSvc interfaces.IProduct) *ProductController {
	return &ProductController{ProductService: productSvc}
}

func (p ProductController) InsertProduct(c *gin.Context) {
	fmt.Println("Invoked controller")
	var product entities.Product
	err := c.BindJSON(&product)
	if err != nil {
		return
	}
	result, err := p.ProductService.Insert(&product)
	if err != nil {
		return
	} else {
		c.IndentedJSON(http.StatusCreated, result)
	}

	//extract the product from the request Object
	//call the service.insert
	//p.ProductService.InsertOne()
}

func (p ProductController) GetProducts(c *gin.Context) {
	result, err := p.ProductService.GetProducts()
	if err != nil {
		return
	} else {
		c.IndentedJSON(http.StatusCreated, result)
	}
}
