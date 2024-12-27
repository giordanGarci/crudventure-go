package main

import (
	"github.com/gin-gonic/gin"
	"github.com/giordanGarci/crudventure-go/controller"
	"github.com/giordanGarci/crudventure-go/db"
	"github.com/giordanGarci/crudventure-go/repository"
	"github.com/giordanGarci/crudventure-go/usecase"
)

func main() {
	server := gin.Default()

	dbConn, err := db.ConnectDB() // camada de db
	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(dbConn)         // camada de repository
	ProductUseCase := usecase.NewProductUsecase(ProductRepository)       // camada de usecase
	ProductController := controller.NewProductController(ProductUseCase) // camada de controller

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)

	server.Run(":8000")

}
