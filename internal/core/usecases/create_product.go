package usecases

import (
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/errors"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/repositories"
)

type (
	CreateProductUseCase interface {
		Execute(request CreateProductRequest) (*entities.Product, error)
	}

	CreateProductRequest struct {
		Name        string
		Description string
		Price       int
		Discount    int
	}

	createProductUseCase struct {
		productsRepository repositories.ProductsRepository
	}
)

func NewCreateProductUseCase(productsRepository repositories.ProductsRepository) CreateProductUseCase {
	return &createProductUseCase{
		productsRepository: productsRepository,
	}
}

func (uc createProductUseCase) Execute(request CreateProductRequest) (*entities.Product, error) {
	product, err := uc.productsRepository.FindByName(request.Name)
	if err != nil {
		return nil, err
	}

	if product != nil {
		return nil, errors.NewProductAlreadyExists()
	}

	product = entities.NewProduct(request.Name, request.Price, entities.ProductOptions{
		Description: request.Description,
		Discount:    request.Discount,
	})

	err = uc.productsRepository.Create(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}
