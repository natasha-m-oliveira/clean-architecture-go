package usecases

import (
	"testing"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/errors"
	"github.com/natasha-m-oliveira/clean-architecture-go/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetCartById(t *testing.T) {

	type TestGetCartByIdSetup struct {
		cartsRepository    *mocks.MockCartsRepository
		getCartByIdUseCase GetCartByIdUseCase
	}

	setup := func() TestGetCartByIdSetup {
		cartsRepository := new(mocks.MockCartsRepository)
		getCartByIdUseCase := NewGetCartByIdUseCase(cartsRepository)

		return TestGetCartByIdSetup{
			cartsRepository:    cartsRepository,
			getCartByIdUseCase: getCartByIdUseCase,
		}
	}

	t.Run("should be able to get cart by id", func(t *testing.T) {
		s := setup()

		s.cartsRepository.On("FindById", "any-cart-id").Return(&entities.Cart{
			Id: "any-cart-id",
		}, nil)

		getCartByIdResponse, err := s.getCartByIdUseCase.Execute(GetCartByIdRequest{
			Id: "any-cart-id",
		})

		assert.NoError(t, err)
		assert.Equal(t, "any-cart-id", getCartByIdResponse.Cart.Id)
	})

	t.Run("should not be able to get non-existent cart by id", func(t *testing.T) {
		s := setup()

		expectedError := errors.NewCartNotFound()

		s.cartsRepository.On("FindById", "invalid-cart-id").Return(nil, expectedError)

		_, err := s.getCartByIdUseCase.Execute(GetCartByIdRequest{
			Id: "invalid-cart-id",
		})

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, expectedError.Error(), err.Error())
	})

	t.Run("should not be able to get cart by id if repository returns an error when finding the cart", func(t *testing.T) {
		s := setup()

		expectedError := errors.NewInternalServerError()

		s.cartsRepository.On("FindById", "any-cart-id").Return(nil, expectedError)

		_, err := s.getCartByIdUseCase.Execute(GetCartByIdRequest{
			Id: "any-cart-id",
		})

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, expectedError.Error(), err.Error())
	})
}
