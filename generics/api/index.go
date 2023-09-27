package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// Define a generic resource type with a type parameter T.
type Resource[T any] struct {
    Data []T
    NonCollection T
}
var customerDetails = Resource[string]{
    Data: []string{"Alice", "Bob", "Charlie"},
}

  
var userResource = Resource[string]{
    Data: []string{"Alice", "Bob", "Charlie"},
}

var productResource = Resource[map[string]interface {}]{
    Data: []map[string]interface {}{
        {"id": 1, "name": "Product A", "price": 10.99},
        {"id": 2, "name": "Product B", "price": 20.99},
    },
}

func GetAll[T any](c *gin.Context) {
    resource := c.MustGet("resource").(Resource[T])
    c.JSON(http.StatusOK, resource.Data)
}

func main() {
    r := gin.Default()

    // Handle user resource
    r.GET("/users", func(c *gin.Context) {
        c.Set("resource", userResource)
        GetAll[string](c)
    })

    // Handle product resource
    r.GET("/products", func(c *gin.Context) {
        c.Set("resource", productResource)
        GetAll[map[string]interface {}](c)
    })
    r.GET("/profiles", func(c *gin.Context) {
        c.Set("resource", customerDetails)
        GetAll[map[string]interface {}](c)
    })

    r.Run(":8080")
}
