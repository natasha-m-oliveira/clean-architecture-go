package products

import (
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/errors"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/repositories"
)

type CreateProductRequest struct {
	Name        string
	Description string
	Price       int
	Discount    int
}

type CreateProductResponse struct {
	Product entities.Product
}

type CreateProductUseCase struct {
	productsRepository repositories.ProductsRepository
}

func NewCreateProductUseCase(productsRepository repositories.ProductsRepository) *CreateProductUseCase {
	return &CreateProductUseCase{
		productsRepository: productsRepository,
	}
}

func (uc *CreateProductUseCase) Execute(request CreateProductRequest) (*CreateProductResponse, error) {
	productAlreadyExists, err := uc.productsRepository.FindByName(request.Name)
	if err != nil && err.Error() != (&errors.ProductNotFound{}).Error() {
		return nil, err
	}

	if productAlreadyExists != nil {
		return nil, &errors.ProductAlreadyExists{}
	}

	product := entities.NewProduct(request.Name, request.Price, entities.ProductOptions{
		Description: request.Description,
		Discount:    request.Discount,
	})

	err = uc.productsRepository.Create(product)
	if err != nil {
		return nil, err
	}

	return &CreateProductResponse{
		Product: *product,
	}, nil
}
