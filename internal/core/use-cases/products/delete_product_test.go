package products

import (
	"testing"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/errors"
	"github.com/natasha-m-oliveira/clean-architecture-go/test/repositories"
	"github.com/stretchr/testify/assert"
)

func TestDeleteProduct(t *testing.T) {
	var productsRepository repositories.InMemoryProductsRepository
	var createProductUseCase CreateProductUseCase
	var deleteProductUseCase DeleteProductUseCase

	setup := func() {
		productsRepository = *repositories.NewInMemoryProductsRepository()

		createProductUseCase = *NewCreateProductUseCase(&productsRepository)
		deleteProductUseCase = *NewDeleteProductUseCase(&productsRepository)
	}

	t.Run("should be able to delete a product", func(t *testing.T) {
		setup()

		createProductResponse, err := createProductUseCase.Execute(CreateProductRequest{
			Name:  "Malta",
			Price: 328621,
		})

		assert.NoError(t, err)

		err = deleteProductUseCase.Execute(DeleteProductRequest{
			Id: createProductResponse.Product.Id,
		})

		assert.NoError(t, err)
	})

	t.Run("should not be able to delete a product that does not exist", func(t *testing.T) {
		setup()

		err := deleteProductUseCase.Execute(DeleteProductRequest{
			Id: "invalid-product-id",
		})

		assert.Error(t, err)
		assert.Equal(t, (&errors.ProductNotFound{}).Error(), err.Error())
	})
}
