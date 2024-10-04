package usecases

import (
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/repositories"
)

type (
	ListProductsUseCase interface {
		Execute(request ListProductsRequest) (*ListProductsResponse, error)
	}

	ListProductsRequest struct {
	}

	ListProductsResponse struct {
		Products []entities.Product
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

func (uc listProductsUseCase) Execute(request ListProductsRequest) (*ListProductsResponse, error) {
	products, err := uc.productsRepository.List()

	return &ListProductsResponse{
		Products: products,
	}, err
}
