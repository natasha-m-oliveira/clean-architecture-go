package usecases

import (
	"testing"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/errors"
	"github.com/natasha-m-oliveira/clean-architecture-go/test/repositories"
	"github.com/stretchr/testify/assert"
)

func TestUpdateCartItems(t *testing.T) {
	var cartsRepository repositories.InMemoryCartsRepository
	var productsRepository repositories.InMemoryProductsRepository
	var createProductUseCase CreateProductUseCase
	var createCartUseCase CreateCartUseCase
	var updateCartItemsUseCase UpdateCartItemsUseCase

	setup := func() {
		cartsRepository = *repositories.NewInMemoryCartsRepository()
		productsRepository = *repositories.NewInMemoryProductsRepository()

		createProductUseCase = NewCreateProductUseCase(&productsRepository)
		createCartUseCase = NewCreateCartUseCase(&cartsRepository, &productsRepository)
		updateCartItemsUseCase = NewUpdateCartItemsUseCase(&cartsRepository, &productsRepository)
	}

	t.Run("should be able to update items in a cart", func(t *testing.T) {
		setup()

		createProductResponse, err := createProductUseCase.Execute(CreateProductRequest{
			Name:  "Nauru",
			Price: 172025,
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

		updateCartItemsResponse, err := updateCartItemsUseCase.Execute(UpdateCartItemsRequest{
			Id: createCartResponse.Cart.Id,
			Items: []struct {
				ProductId string
				Quantity  int
			}{
				{
					ProductId: createProductResponse.Product.Id,
					Quantity:  1,
				},
			},
		})

		assert.NoError(t, err)
		assert.Equal(t, 1, updateCartItemsResponse.Cart.Items[0].Quantity)
	})

	t.Run("should not be able to update items in a non-existent cart", func(t *testing.T) {
		setup()

		_, err := updateCartItemsUseCase.Execute(UpdateCartItemsRequest{
			Id: "invalid-cart-id",
		})

		assert.Error(t, err)
		assert.Equal(t, (&errors.CartNotFound{}).Error(), err.Error())
	})

	t.Run("should not be able to update cart items if one of the products does not exist", func(t *testing.T) {
		setup()

		createProductResponse, err := createProductUseCase.Execute(CreateProductRequest{
			Name:  "Botswana",
			Price: 69308,
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

		_, err = updateCartItemsUseCase.Execute(UpdateCartItemsRequest{
			Id: createCartResponse.Cart.Id,
			Items: []struct {
				ProductId string
				Quantity  int
			}{
				{
					ProductId: createProductResponse.Product.Id,
					Quantity:  2,
				},
				{
					ProductId: "invalid-product-id",
					Quantity:  1,
				},
			},
		})

		assert.Error(t, err)
		assert.Equal(t, (&errors.ProductNotFound{}).Error(), err.Error())
	})
}
