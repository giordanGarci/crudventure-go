package usecase

type productUsecase struct {
	// repository
}

func NewProductUsecase() *productUsecase {
	return &productUsecase{}
}

func (pu *productUsecase) GetProducts() {
	// repository.GetProducts()
}
