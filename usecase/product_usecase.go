package usecase

import (
	"github.com/giordanGarci/crudventure-go/model"
	"github.com/giordanGarci/crudventure-go/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId
	return product, nil
}

func (pu *ProductUsecase) GetProductById(productId int64) (*model.Product, error) {
	product, err := pu.repository.GetProductById(productId)
	if err != nil {
		return nil, err
	}
	return product, nil
}
