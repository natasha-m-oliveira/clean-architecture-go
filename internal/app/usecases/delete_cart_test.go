package usecases

import (
	"testing"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/errors"
	"github.com/natasha-m-oliveira/clean-architecture-go/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestDeleteCart(t *testing.T) {

	type TestDeleteCartSetup struct {
		cartsRepository   *mocks.MockCartsRepository
		deleteCartUseCase DeleteCartUseCase
	}

	setup := func() TestDeleteCartSetup {
		cartsRepository := new(mocks.MockCartsRepository)

		deleteCartUseCase := NewDeleteCartUseCase(cartsRepository)

		return TestDeleteCartSetup{
			cartsRepository:   cartsRepository,
			deleteCartUseCase: deleteCartUseCase,
		}
	}

	t.Run("should be able to delete a cart", func(t *testing.T) {
		s := setup()

		s.cartsRepository.On("FindById", "any-cart-id").Return(&entities.Cart{
			Id: "any-cart-id",
		}, nil)
		s.cartsRepository.On("DeleteById", "any-cart-id").Return(nil)

		err := s.deleteCartUseCase.Execute(DeleteCartRequest{
			Id: "any-cart-id",
		})

		assert.NoError(t, err)
	})

	t.Run("should not be able to delete a non-existent cart", func(t *testing.T) {
		s := setup()

		expectedError := errors.NewCartNotFound()

		s.cartsRepository.On("FindById", "invalid-cart-id").Return(nil, expectedError)

		err := s.deleteCartUseCase.Execute(DeleteCartRequest{
			Id: "invalid-cart-id",
		})

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, expectedError.Error(), err.Error())
	})

	t.Run("should not be able to delete a cart if the repository returns an error when finding the cart", func(t *testing.T) {
		s := setup()

		expectedError := errors.NewCartNotFound()

		s.cartsRepository.On("FindById", "any-cart-id").Return(nil, expectedError)

		err := s.deleteCartUseCase.Execute(DeleteCartRequest{
			Id: "any-cart-id",
		})

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, expectedError.Error(), err.Error())
	})

	t.Run("should not be able to delete a cart if the repository returns an error when deleting the cart", func(t *testing.T) {
		s := setup()

		expectedError := errors.NewCartNotFound()

		s.cartsRepository.On("FindById", "any-cart-id").Return(&entities.Cart{
			Id: "any-cart-id",
		}, nil)
		s.cartsRepository.On("DeleteById", "any-cart-id").Return(expectedError)

		err := s.deleteCartUseCase.Execute(DeleteCartRequest{
			Id: "any-cart-id",
		})

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, expectedError.Error(), err.Error())
	})
}
