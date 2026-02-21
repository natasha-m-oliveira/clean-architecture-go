package usecases

import (
	"testing"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/errors"
	"github.com/natasha-m-oliveira/clean-architecture-go/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestListProducts(t *testing.T) {
	type TestListProductsSetup struct {
		productsRepository  *mocks.MockProductsRepository
		listProductsUseCase ListProductsUseCase
	}

	setup := func() TestListProductsSetup {
		productsRepository := new(mocks.MockProductsRepository)
		listProductsUseCase := NewListProductsUseCase(productsRepository)

		return TestListProductsSetup{
			productsRepository:  productsRepository,
			listProductsUseCase: listProductsUseCase,
		}
	}

	t.Run("should be able to list products", func(t *testing.T) {
		s := setup()

		s.productsRepository.On("FindAll").Return([]entities.Product{
			{
				Id:          "any-product-id",
				Name:        "any-product-name",
				Description: "any-product-description",
				Price:       10.0,
			},
		}, nil)

		listProductsResponse, err := s.listProductsUseCase.Execute(ListProductsRequest{})

		assert.NoError(t, err)
		assert.Equal(t, 1, len(listProductsResponse.Products))
		assert.Equal(t, "any-product-id", listProductsResponse.Products[0].Id)
	})

	t.Run("should return error if repository returns error", func(t *testing.T) {
		s := setup()

		expectedError := errors.NewInternalServerError()

		s.productsRepository.On("FindAll").Return([]entities.Product{}, expectedError)

		listProductsResponse, err := s.listProductsUseCase.Execute(ListProductsRequest{})

		assert.NotNil(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, expectedError.Error(), err.Error())
		assert.Nil(t, listProductsResponse)
	})
}
