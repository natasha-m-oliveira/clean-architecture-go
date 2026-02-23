package usecases

import (
	"testing"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/errors"
	"github.com/natasha-m-oliveira/clean-architecture-go/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCart(t *testing.T) {

	type TestCreateCartSetup struct {
		cartsRepository    *mocks.MockCartsRepository
		productsRepository *mocks.MockProductsRepository
		createCartUseCase  CreateCartUseCase
	}

	setup := func() TestCreateCartSetup {
		cartsRepository := new(mocks.MockCartsRepository)
		productsRepository := new(mocks.MockProductsRepository)

		createCartUseCase := NewCreateCartUseCase(cartsRepository, productsRepository)

		return TestCreateCartSetup{
			cartsRepository:    cartsRepository,
			productsRepository: productsRepository,
			createCartUseCase:  createCartUseCase,
		}
	}

	t.Run("should be able to create a new cart", func(t *testing.T) {
		s := setup()

		s.productsRepository.On("FindById", "any-product-id").Return(&entities.Product{
			Id:    "any-product-id",
			Name:  "Any Product",
			Price: 1000,
		}, nil)

		s.cartsRepository.On("Create", mock.AnythingOfType("*entities.Cart")).Return(nil)

		createCartResponse, err := s.createCartUseCase.Execute(CreateCartRequest{
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

		assert.NoError(t, err)
		assert.Equal(t, 1, len(createCartResponse.Cart.Items))
		s.cartsRepository.AssertExpectations(t)
		s.productsRepository.AssertExpectations(t)
	})

	t.Run("should not be able to create a cart if one of the products does not exist", func(t *testing.T) {
		s := setup()

		expectedError := errors.NewProductNotFound()

		s.productsRepository.On("FindById", "invalid-product-id").Return(nil, expectedError)

		_, err := s.createCartUseCase.Execute(CreateCartRequest{
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
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, expectedError.Error(), err.Error())
	})

	t.Run("should not be able to create a cart if repository returns an error when checking if product exists", func(t *testing.T) {
		s := setup()

		expectedError := errors.NewInternalServerError()

		s.productsRepository.On("FindById", "any-product-id").Return(nil, expectedError)

		_, err := s.createCartUseCase.Execute(CreateCartRequest{
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

	t.Run("should not be able to create a cart if repository returns an error when creating the cart", func(t *testing.T) {
		s := setup()

		s.productsRepository.On("FindById", "any-product-id").Return(&entities.Product{
			Id:    "any-product-id",
			Name:  "Any Product",
			Price: 1000,
		}, nil)

		expectedError := errors.NewInternalServerError()

		s.cartsRepository.On("Create", mock.AnythingOfType("*entities.Cart")).Return(expectedError)

		_, err := s.createCartUseCase.Execute(CreateCartRequest{
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
		s.productsRepository.AssertExpectations(t)
		s.cartsRepository.AssertExpectations(t)
	})
}
