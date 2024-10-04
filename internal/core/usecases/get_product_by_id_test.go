package usecases

import (
	"testing"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/errors"
	"github.com/natasha-m-oliveira/clean-architecture-go/test/repositories"
	"github.com/stretchr/testify/assert"
)

func TestGetProductById(t *testing.T) {
	var productsRepository repositories.InMemoryProductsRepository
	var createProductUseCase CreateProductUseCase
	var getProductByIdUseCase GetProductByIdUseCase

	setup := func() {
		productsRepository = *repositories.NewInMemoryProductsRepository()

		createProductUseCase = NewCreateProductUseCase(&productsRepository)
		getProductByIdUseCase = NewGetProductByIdUseCase(&productsRepository)
	}

	t.Run("should be able to get product by id", func(t *testing.T) {
		setup()

		createProductResponse, err := createProductUseCase.Execute(CreateProductRequest{
			Name:  "El Salvador",
			Price: 34512,
		})

		assert.NoError(t, err)

		getProductByIdResponse, err := getProductByIdUseCase.Execute(GetProductByIdRequest{
			Id: createProductResponse.Product.Id,
		})

		assert.NoError(t, err)
		assert.Equal(t, createProductResponse.Product.Id, getProductByIdResponse.Product.Id)
	})

	t.Run("should not be able to get non-existent product by id", func(t *testing.T) {
		setup()

		_, err := getProductByIdUseCase.Execute(GetProductByIdRequest{
			Id: "invalid-product-id",
		})

		assert.Error(t, err)
		assert.Equal(t, (&errors.ProductNotFound{}).Error(), err.Error())
	})
}
