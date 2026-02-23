package usecases

import (
	"testing"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/errors"
	"github.com/natasha-m-oliveira/clean-architecture-go/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateProduct(t *testing.T) {

	type TestCreateProductSetup struct {
		productsRepository   *mocks.MockProductsRepository
		createProductUseCase CreateProductUseCase
	}

	setup := func() TestCreateProductSetup {
		productsRepository := new(mocks.MockProductsRepository)

		createProductUseCase := NewCreateProductUseCase(productsRepository)

		return TestCreateProductSetup{
			productsRepository:   productsRepository,
			createProductUseCase: createProductUseCase,
		}
	}

	t.Run("should be able to create a new product", func(t *testing.T) {
		s := setup()

		s.productsRepository.On("FindByName", "Suriname").Return(nil, nil)
		s.productsRepository.On("Create", mock.AnythingOfType("*entities.Product")).Return(nil)

		response, err := s.createProductUseCase.Execute(CreateProductRequest{
			Name:  "Suriname",
			Price: 780643,
		})

		assert.NoError(t, err)
		assert.NotEmpty(t, response.Product.Id)
		assert.Equal(t, "Suriname", response.Product.Name)
		s.productsRepository.AssertExpectations(t)
	})

	t.Run("should not be able to create a product already exists", func(t *testing.T) {
		s := setup()

		expectedError := errors.NewProductAlreadyExists()

		product := &entities.Product{
			Id:    "any-product-id",
			Name:  "Sweden",
			Price: 37516,
		}

		s.productsRepository.On("FindByName", "Sweden").Return(product, nil)

		_, err := s.createProductUseCase.Execute(CreateProductRequest{
			Name:  "Sweden",
			Price: 37516,
		})

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, expectedError.Error(), err.Error())
		s.productsRepository.AssertExpectations(t)
	})

	t.Run("should not be able to create a product if repository returns an error when creating the product", func(t *testing.T) {
		s := setup()

		expectedError := errors.NewInternalServerError()
		s.productsRepository.On("FindByName", "Denmark").Return(nil, nil)
		s.productsRepository.On("Create", mock.AnythingOfType("*entities.Product")).Return(expectedError)

		_, err := s.createProductUseCase.Execute(CreateProductRequest{
			Name:  "Denmark",
			Price: 183654,
		})

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, expectedError.Error(), err.Error())
		s.productsRepository.AssertExpectations(t)
	})

	t.Run("should not be able to create a product if repository returns an error when checking if product already exists", func(t *testing.T) {
		s := setup()

		expectedError := errors.NewInternalServerError()
		s.productsRepository.On("FindByName", "Finland").Return(nil, expectedError)

		_, err := s.createProductUseCase.Execute(CreateProductRequest{
			Name:  "Finland",
			Price: 246813,
		})

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, expectedError.Error(), err.Error())
		s.productsRepository.AssertExpectations(t)
	})
}
