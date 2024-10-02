package products

import (
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/repositories"
)

type ListProductsRequest struct {
}

type ListProductsResponse struct {
	Products []entities.Product
}

type ListProductsUseCase struct {
	productsRepository repositories.ProductsRepository
}

func NewListProductsUseCase(productsRepository repositories.ProductsRepository) *ListProductsUseCase {
	return &ListProductsUseCase{
		productsRepository: productsRepository,
	}
}

func (uc *ListProductsUseCase) Execute(request ListProductsRequest) (*ListProductsResponse, error) {
	products, err := uc.productsRepository.List()

	return &ListProductsResponse{
		Products: products,
	}, err
}
