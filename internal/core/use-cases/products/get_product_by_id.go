package products

import (
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/repositories"
)

type GetProductByIdRequest struct {
	Id string
}

type GetProductByIdResponse struct {
	Product entities.Product
}

type GetProductByIdUseCase struct {
	productsRepository repositories.ProductsRepository
}

func NewGetProductByIdUseCase(productsRepository repositories.ProductsRepository) *GetProductByIdUseCase {
	return &GetProductByIdUseCase{
		productsRepository: productsRepository,
	}
}

func (uc *GetProductByIdUseCase) Execute(request GetProductByIdRequest) (*GetProductByIdResponse, error) {
	product, err := uc.productsRepository.FindById(request.Id)
	if err != nil {
		return nil, err
	}

	return &GetProductByIdResponse{
		Product: *product,
	}, nil
}
