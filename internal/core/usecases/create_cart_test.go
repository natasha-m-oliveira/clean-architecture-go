package usecases

import (
	"testing"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/errors"
	"github.com/natasha-m-oliveira/clean-architecture-go/test/repositories"
	"github.com/stretchr/testify/assert"
)

func TestCreateCart(t *testing.T) {
	var cartsRepository repositories.InMemoryCartsRepository
	var productsRepository repositories.InMemoryProductsRepository
	var createProductUseCase CreateProductUseCase
	var createCartUseCase CreateCartUseCase

	setup := func() {
		cartsRepository = *repositories.NewInMemoryCartsRepository()
		productsRepository = *repositories.NewInMemoryProductsRepository()

		createProductUseCase = NewCreateProductUseCase(&productsRepository)
		createCartUseCase = NewCreateCartUseCase(&cartsRepository, &productsRepository)
	}

	t.Run("should be able to create a new cart", func(t *testing.T) {
		setup()

		createProductResponse, err := createProductUseCase.Execute(CreateProductRequest{
			Name:  "Malta",
			Price: 328621,
		})

		assert.NoError(t, err)

		createCartResponse, err := createCartUseCase.Execute(CreateCartRequest{
			Items: []struct {
				ProductId string
				Quantity  int
			}{
				{
					ProductId: createProductResponse.Product.Id,
					Quantity:  2,
				},
			},
		})

		assert.NoError(t, err)
		assert.Equal(t, 1, len(createCartResponse.Cart.Items))
	})

	t.Run("should not be able to create a cart if one of the products does not exist", func(t *testing.T) {
		setup()

		_, err := createCartUseCase.Execute(CreateCartRequest{
			Items: []struct {
				ProductId string
				Quantity  int
			}{
				{
					ProductId: "invalid-product-id",
					Quantity:  2,
				},
			},
		})

		assert.Error(t, err)
		assert.Equal(t, (&errors.ProductNotFound{}).Error(), err.Error())
	})
}
