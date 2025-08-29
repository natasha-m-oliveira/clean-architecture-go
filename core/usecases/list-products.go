package usecases

import (
	"github.com/natasha-m-oliveira/clean-architecture-go/core/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/core/repositories"
)

type (
	ListProductsUseCase interface {
		Execute(request ListProductsRequest) ([]entities.Product, error)
	}

	ListProductsRequest struct {
	}

	listProductsUseCase struct {
		productsRepository repositories.ProductsRepository
	}
)

func NewListProductsUseCase(productsRepository repositories.ProductsRepository) ListProductsUseCase {
	return &listProductsUseCase{
		productsRepository: productsRepository,
	}
}

func (uc listProductsUseCase) Execute(request ListProductsRequest) ([]entities.Product, error) {
	return uc.productsRepository.List()
}
