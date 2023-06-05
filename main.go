package main

import (
	"go-with-gin/controllers"
	"go-with-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", controllers.Index)
	r.GET("/api/product/:id", controllers.GetProductById)
	r.POST("/api/product", controllers.CreateProduct)
	r.PUT("/api/product/:id", controllers.UpdateProduct)
	r.DELETE("/api/product", controllers.DeleteProduct)

	r.Run()
}
