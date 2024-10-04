package usecases

import (
	"testing"

	"github.com/natasha-m-oliveira/clean-architecture-go/test/repositories"
	"github.com/stretchr/testify/assert"
)

func TestListProducts(t *testing.T) {
	var productsRepository repositories.InMemoryProductsRepository
	var createProductUseCase CreateProductUseCase
	var listProductsUseCase ListProductsUseCase

	setup := func() {
		productsRepository = *repositories.NewInMemoryProductsRepository()

		createProductUseCase = NewCreateProductUseCase(&productsRepository)
		listProductsUseCase = NewListProductsUseCase(&productsRepository)
	}

	t.Run("should be able to list products", func(t *testing.T) {
		setup()

		createProductResponse, err := createProductUseCase.Execute(CreateProductRequest{
			Name:  "Mauritania",
			Price: 3594,
		})

		assert.NoError(t, err)

		listProductsResponse, err := listProductsUseCase.Execute(ListProductsRequest{})

		assert.NoError(t, err)
		assert.Equal(t, 1, len(listProductsResponse.Products))
		assert.Equal(t, createProductResponse.Product.Id, listProductsResponse.Products[0].Id)
	})
}
