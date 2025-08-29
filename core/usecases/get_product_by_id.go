package usecases

import (
	"github.com/natasha-m-oliveira/clean-architecture-go/core/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/core/errors"
	"github.com/natasha-m-oliveira/clean-architecture-go/core/repositories"
)

type (
	GetProductByIdUseCase interface {
		Execute(request GetProductByIdRequest) (*entities.Product, error)
	}

	GetProductByIdRequest struct {
		Id string
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

func (uc getProductByIdUseCase) Execute(request GetProductByIdRequest) (*entities.Product, error) {
	product, err := uc.productsRepository.FindById(request.Id)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, errors.NewProductNotFound()
	}

	return product, nil
}
