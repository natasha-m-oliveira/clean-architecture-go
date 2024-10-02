package products

import "github.com/natasha-m-oliveira/clean-architecture-go/internal/core/repositories"

type DeleteProductRequest struct {
	Id string
}

type DeleteProductUseCase struct {
	productsRepository repositories.ProductsRepository
}

func NewDeleteProductUseCase(productsRepository repositories.ProductsRepository) *DeleteProductUseCase {
	return &DeleteProductUseCase{
		productsRepository: productsRepository,
	}
}

func (uc *DeleteProductUseCase) Execute(request DeleteProductRequest) error {
	_, err := uc.productsRepository.FindById(request.Id)
	if err != nil {
		return err
	}

	err = uc.productsRepository.DeleteById(request.Id)

	return err
}
