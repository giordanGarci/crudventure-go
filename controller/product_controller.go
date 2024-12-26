package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/giordanGarci/crudventure-go/model"
)

type productController struct {
	// Usecase
}

func NewProductController() *productController {
	return &productController{}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		{
			ID:    1,
			Name:  "Batata Frita",
			Price: 20,
		},
	}
	ctx.JSON(http.StatusOK, products)
}
