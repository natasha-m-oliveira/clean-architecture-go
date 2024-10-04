package usecases

import (
	"testing"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/errors"
	"github.com/natasha-m-oliveira/clean-architecture-go/test/repositories"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	var productsRepository repositories.InMemoryProductsRepository
	var createProductUseCase CreateProductUseCase

	setup := func() {
		productsRepository = *repositories.NewInMemoryProductsRepository()

		createProductUseCase = NewCreateProductUseCase(&productsRepository)
	}

	t.Run("should be able to create a new product", func(t *testing.T) {
		setup()

		response, err := createProductUseCase.Execute(CreateProductRequest{
			Name:  "Suriname",
			Price: 780643,
		})

		assert.NoError(t, err)
		assert.NotEmpty(t, response.Product.Id)
		assert.Equal(t, "Suriname", response.Product.Name)
	})

	t.Run("should not be able to create a product already exists", func(t *testing.T) {
		setup()

		_, err := createProductUseCase.Execute(CreateProductRequest{
			Name:  "Sweden",
			Price: 37516,
		})

		assert.NoError(t, err)

		_, err = createProductUseCase.Execute(CreateProductRequest{
			Name:  "Sweden",
			Price: 37516,
		})

		assert.Error(t, err)
		assert.Equal(t, (&errors.ProductAlreadyExists{}).Error(), err.Error())
	})
}
