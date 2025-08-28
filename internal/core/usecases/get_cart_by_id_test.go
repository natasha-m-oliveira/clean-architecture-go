package usecases

import (
	"testing"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/errors"
	"github.com/natasha-m-oliveira/clean-architecture-go/test/repositories"
	"github.com/stretchr/testify/assert"
)

func TestGetCartById(t *testing.T) {
	var cartsRepository repositories.InMemoryCartsRepository
	var productsRepository repositories.InMemoryProductsRepository
	var createProductUseCase CreateProductUseCase
	var createCartUseCase CreateCartUseCase
	var getCartByIdUseCase GetCartByIdUseCase

	setup := func() {
		cartsRepository = *repositories.NewInMemoryCartsRepository()
		productsRepository = *repositories.NewInMemoryProductsRepository()

		createProductUseCase = NewCreateProductUseCase(&productsRepository)
		createCartUseCase = NewCreateCartUseCase(&cartsRepository, &productsRepository)
		getCartByIdUseCase = NewGetCartByIdUseCase(&cartsRepository)
	}

	t.Run("should be able to get cart by id", func(t *testing.T) {
		setup()

		createProductResponse, err := createProductUseCase.Execute(CreateProductRequest{
			Name:  "Guernsey",
			Price: 1893,
		})

		assert.NoError(t, err)

		createCartResponse, err := createCartUseCase.Execute(CreateCartRequest{
			Items: []CreateCartItemRequest{
				{
					ProductId: createProductResponse.Id,
					Quantity:  2,
				},
			},
		})

		assert.NoError(t, err)

		getCartByIdResponse, err := getCartByIdUseCase.Execute(GetCartByIdRequest{
			Id: createCartResponse.Id,
		})

		assert.NoError(t, err)
		assert.Equal(t, createCartResponse.Id, getCartByIdResponse.Id)
	})

	t.Run("should not be able to get non-existent cart by id", func(t *testing.T) {
		setup()

		_, err := getCartByIdUseCase.Execute(GetCartByIdRequest{
			Id: "invalid-cart-id",
		})

		assert.Error(t, err)
		assert.Equal(t, errors.NewCartNotFound(), err)
	})
}
