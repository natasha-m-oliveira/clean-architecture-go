package usecases

import (
	"testing"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/errors"
	"github.com/natasha-m-oliveira/clean-architecture-go/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestDeleteProduct(t *testing.T) {

	type TestDeleteProductSetup struct {
		productsRepository   *mocks.MockProductsRepository
		deleteProductUseCase DeleteProductUseCase
	}

	setup := func() TestDeleteProductSetup {
		productsRepository := new(mocks.MockProductsRepository)

		deleteProductUseCase := NewDeleteProductUseCase(productsRepository)

		return TestDeleteProductSetup{
			productsRepository:   productsRepository,
			deleteProductUseCase: deleteProductUseCase,
		}
	}

	t.Run("should be able to delete a product", func(t *testing.T) {
		s := setup()

		s.productsRepository.On("FindById", "any-product-id").Return(&entities.Product{
			Id: "any-product-id",
		}, nil)
		s.productsRepository.On("DeleteById", "any-product-id").Return(nil)

		err := s.deleteProductUseCase.Execute(DeleteProductRequest{
			Id: "any-product-id",
		})

		assert.NoError(t, err)
	})

	t.Run("should not be able to delete a product that does not exist", func(t *testing.T) {
		s := setup()

		expectedError := errors.NewProductNotFound()

		s.productsRepository.On("FindById", "invalid-product-id").Return(nil, expectedError)

		err := s.deleteProductUseCase.Execute(DeleteProductRequest{
			Id: "invalid-product-id",
		})

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, expectedError.Error(), err.Error())
	})

	t.Run("should not be able to delete a product if repository returns an error when finding the product", func(t *testing.T) {
		s := setup()

		expectedError := errors.NewProductNotFound()

		s.productsRepository.On("FindById", "any-product-id").Return(nil, expectedError)

		err := s.deleteProductUseCase.Execute(DeleteProductRequest{
			Id: "any-product-id",
		})

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, expectedError.Error(), err.Error())
	})

	t.Run("should not be able to delete a product if repository returns an error when deleting the product", func(t *testing.T) {
		s := setup()

		expectedError := errors.NewProductNotFound()

		s.productsRepository.On("FindById", "any-product-id").Return(&entities.Product{
			Id: "any-product-id",
		}, nil)
		s.productsRepository.On("DeleteById", "any-product-id").Return(expectedError)

		err := s.deleteProductUseCase.Execute(DeleteProductRequest{
			Id: "any-product-id",
		})

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, expectedError.Error(), err.Error())
	})
}
