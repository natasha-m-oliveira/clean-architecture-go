package usecases

import (
	"testing"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/errors"
	"github.com/natasha-m-oliveira/clean-architecture-go/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetProductById(t *testing.T) {

	type TestGetProductByIdSetup struct {
		productsRepository    *mocks.MockProductsRepository
		getProductByIdUseCase GetProductByIdUseCase
	}

	setup := func() TestGetProductByIdSetup {
		productsRepository := new(mocks.MockProductsRepository)
		getProductByIdUseCase := NewGetProductByIdUseCase(productsRepository)

		return TestGetProductByIdSetup{
			productsRepository:    productsRepository,
			getProductByIdUseCase: getProductByIdUseCase,
		}
	}

	t.Run("should be able to get product by id", func(t *testing.T) {
		s := setup()

		s.productsRepository.On("FindById", "any-product-id").Return(&entities.Product{
			Id: "any-product-id",
		}, nil)

		getProductByIdResponse, err := s.getProductByIdUseCase.Execute(GetProductByIdRequest{
			Id: "any-product-id",
		})

		assert.NoError(t, err)
		assert.Equal(t, "any-product-id", getProductByIdResponse.Product.Id)
	})

	t.Run("should not be able to get non-existent product by id", func(t *testing.T) {
		s := setup()

		expectedError := errors.NewProductNotFound()

		s.productsRepository.On("FindById", "invalid-product-id").Return(nil, expectedError)

		_, err := s.getProductByIdUseCase.Execute(GetProductByIdRequest{
			Id: "invalid-product-id",
		})

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, expectedError.Error(), err.Error())
	})

	t.Run("should not be able to get product by id if repository returns an error when finding the product", func(t *testing.T) {
		s := setup()

		expectedError := errors.NewInternalServerError()

		s.productsRepository.On("FindById", "any-product-id").Return(nil, expectedError)

		_, err := s.getProductByIdUseCase.Execute(GetProductByIdRequest{
			Id: "any-product-id",
		})

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, expectedError.Error(), err.Error())
	})
}
