package usecases

import "github.com/natasha-m-oliveira/clean-architecture-go/internal/core/repositories"

type (
	DeleteProductUseCase interface {
		Execute(request DeleteProductRequest) error
	}

	DeleteProductRequest struct {
		Id string
	}

	deleteProductUseCase struct {
		productsRepository repositories.ProductsRepository
	}
)

func NewDeleteProductUseCase(productsRepository repositories.ProductsRepository) DeleteProductUseCase {
	return &deleteProductUseCase{
		productsRepository: productsRepository,
	}
}

func (uc deleteProductUseCase) Execute(request DeleteProductRequest) error {
	_, err := uc.productsRepository.FindById(request.Id)
	if err != nil {
		return err
	}

	err = uc.productsRepository.DeleteById(request.Id)

	return err
}
