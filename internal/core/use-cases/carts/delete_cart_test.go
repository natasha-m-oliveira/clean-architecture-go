package carts

import (
	"testing"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/errors"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/use-cases/products"
	"github.com/natasha-m-oliveira/clean-architecture-go/test/repositories"
	"github.com/stretchr/testify/assert"
)

func TestDeleteCart(t *testing.T) {
	var cartsRepository repositories.InMemoryCartsRepository
	var productsRepository repositories.InMemoryProductsRepository
	var createProductUseCase products.CreateProductUseCase
	var createCartUseCase CreateCartUseCase
	var deleteCartUseCase DeleteCartUseCase

	setup := func() {
		cartsRepository = *repositories.NewInMemoryCartsRepository()
		productsRepository = *repositories.NewInMemoryProductsRepository()

		createProductUseCase = *products.NewCreateProductUseCase(&productsRepository)
		createCartUseCase = *NewCreateCartUseCase(&cartsRepository, &productsRepository)
		deleteCartUseCase = *NewDeleteCartUseCase(&cartsRepository)
	}

	t.Run("should be able to delete a cart", func(t *testing.T) {
		setup()

		createProductResponse, err := createProductUseCase.Execute(products.CreateProductRequest{
			Name:  "Antigua and Barbuda",
			Price: 30479,
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

		err = deleteCartUseCase.Execute(DeleteCartRequest{
			Id: createCartResponse.Cart.Id,
		})

		assert.NoError(t, err)
	})

	t.Run("should not be able to delete a non-existent cart", func(t *testing.T) {
		setup()

		err := deleteCartUseCase.Execute(DeleteCartRequest{
			Id: "invalid-cart-id",
		})

		assert.Error(t, err)
		assert.Equal(t, (&errors.CartNotFound{}).Error(), err.Error())
	})
}
