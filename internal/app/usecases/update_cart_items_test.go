package usecases

import (
	"testing"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/errors"
	"github.com/natasha-m-oliveira/clean-architecture-go/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateCartItems(t *testing.T) {

	type TestUpdateCartItemsSetup struct {
		cartsRepository        *mocks.MockCartsRepository
		productsRepository     *mocks.MockProductsRepository
		updateCartItemsUseCase UpdateCartItemsUseCase
	}

	setup := func() TestUpdateCartItemsSetup {
		cartsRepository := new(mocks.MockCartsRepository)
		productsRepository := new(mocks.MockProductsRepository)

		updateCartItemsUseCase := NewUpdateCartItemsUseCase(cartsRepository, productsRepository)

		return TestUpdateCartItemsSetup{
			cartsRepository:        cartsRepository,
			productsRepository:     productsRepository,
			updateCartItemsUseCase: updateCartItemsUseCase,
		}
	}

	t.Run("should be able to update items in a cart", func(t *testing.T) {
		s := setup()

		s.cartsRepository.On("FindById", "any-cart-id").Return(&entities.Cart{
			Id: "any-cart-id",
			Items: []entities.CartItem{
				{
					ProductId: "any-product-id",
					Quantity:  1,
				},
			},
		}, nil)

		s.productsRepository.On("FindById", "any-product-id").Return(&entities.Product{
			Id:    "any-product-id",
			Name:  "any-product-name",
			Price: 10.0,
		}, nil)

		s.cartsRepository.On("Save", mock.AnythingOfType("*entities.Cart")).Return(nil)

		updateCartItemsResponse, err := s.updateCartItemsUseCase.Execute(UpdateCartItemsRequest{
			Id: "any-cart-id",
			Items: []struct {
				ProductId string
				Quantity  int
			}{
				{
					ProductId: "any-product-id",
					Quantity:  1,
				},
			},
		})

		assert.NoError(t, err)
		assert.Equal(t, 1, updateCartItemsResponse.Cart.Items[0].Quantity)
		s.cartsRepository.AssertExpectations(t)
	})

	t.Run("should not be able to update items in a non-existent cart", func(t *testing.T) {
		s := setup()

		expectedError := errors.NewCartNotFound()

		s.cartsRepository.On("FindById", "invalid-cart-id").Return(nil, expectedError)

		_, err := s.updateCartItemsUseCase.Execute(UpdateCartItemsRequest{
			Id: "invalid-cart-id",
		})

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, expectedError.Error(), err.Error())
	})

	t.Run("should not be able to update cart items if one of the products does not exist", func(t *testing.T) {
		s := setup()

		s.cartsRepository.On("FindById", "any-cart-id").Return(&entities.Cart{
			Id: "any-cart-id",
		}, nil)

		s.productsRepository.On("FindById", "any-product-id").Return(&entities.Product{
			Id:    "any-product-id",
			Name:  "any-product-name",
			Price: 10.0,
		}, nil)

		expectedError := errors.NewProductNotFound()

		s.productsRepository.On("FindById", "invalid-product-id").Return(nil, expectedError)

		_, err := s.updateCartItemsUseCase.Execute(UpdateCartItemsRequest{
			Id: "any-cart-id",
			Items: []struct {
				ProductId string
				Quantity  int
			}{
				{
					ProductId: "any-product-id",
					Quantity:  2,
				},
				{
					ProductId: "invalid-product-id",
					Quantity:  1,
				},
			},
		})

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, expectedError.Error(), err.Error())
	})

	t.Run("should not be able to update cart items if repository returns error when finding cart", func(t *testing.T) {
		s := setup()

		expectedError := errors.NewCartNotFound()
		s.cartsRepository.On("FindById", "any-cart-id").Return(nil, expectedError)

		_, err := s.updateCartItemsUseCase.Execute(UpdateCartItemsRequest{
			Id: "any-cart-id",
		})

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, expectedError.Error(), err.Error())
	})

	t.Run("should not be able to update cart items if repository returns error when saving cart", func(t *testing.T) {
		s := setup()

		s.cartsRepository.On("FindById", "any-cart-id").Return(&entities.Cart{
			Id: "any-cart-id",
		}, nil)

		s.productsRepository.On("FindById", "any-product-id").Return(&entities.Product{
			Id:    "any-product-id",
			Name:  "any-product-name",
			Price: 10.0,
		}, nil)

		expectedError := errors.NewInternalServerError()
		s.cartsRepository.On("Save", mock.AnythingOfType("*entities.Cart")).Return(expectedError)

		_, err := s.updateCartItemsUseCase.Execute(UpdateCartItemsRequest{
			Id: "any-cart-id",
			Items: []struct {
				ProductId string
				Quantity  int
			}{
				{
					ProductId: "any-product-id",
					Quantity:  2,
				},
			},
		})

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, expectedError.Error(), err.Error())
	})

	t.Run("should not be able to update cart items if repository returns error when finding product", func(t *testing.T) {
		s := setup()

		s.cartsRepository.On("FindById", "any-cart-id").Return(&entities.Cart{
			Id: "any-cart-id",
		}, nil)

		expectedError := errors.NewInternalServerError()
		s.productsRepository.On("FindById", "any-product-id").Return(nil, expectedError)

		_, err := s.updateCartItemsUseCase.Execute(UpdateCartItemsRequest{
			Id: "any-cart-id",
			Items: []struct {
				ProductId string
				Quantity  int
			}{
				{
					ProductId: "any-product-id",
					Quantity:  2,
				},
			},
		})

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, expectedError.Error(), err.Error())
	})
}
