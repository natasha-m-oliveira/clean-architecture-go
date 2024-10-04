package usecases

import (
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/repositories"
)

type (
	GetProductByIdUseCase interface {
		Execute(request GetProductByIdRequest) (*GetProductByIdResponse, error)
	}

	GetProductByIdRequest struct {
		Id string
	}

	GetProductByIdResponse struct {
		Product entities.Product
	}

	getProductByIdUseCase struct {
		productsRepository repositories.ProductsRepository
	}
)

func NewGetProductByIdUseCase(productsRepository repositories.ProductsRepository) GetProductByIdUseCase {
	return &getProductByIdUseCase{
		productsRepository: productsRepository,
	}
}

func (uc getProductByIdUseCase) Execute(request GetProductByIdRequest) (*GetProductByIdResponse, error) {
	product, err := uc.productsRepository.FindById(request.Id)
	if err != nil {
		return nil, err
	}

	return &GetProductByIdResponse{
		Product: *product,
	}, nil
}
