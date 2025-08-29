package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/natasha-m-oliveira/clean-architecture-go/core/entities"
	usecases_test "github.com/natasha-m-oliveira/clean-architecture-go/test/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListProductsController(t *testing.T) {
	t.Run("should return 200 OK and list of products", func(t *testing.T) {
		mockUseCase := &usecases_test.MockListProductsUseCase{}
		mockProducts := []entities.Product{
			{Id: "1", Name: "Product 1", Price: 100},
			{Id: "2", Name: "Product 2", Price: 200},
		}
		mockUseCase.On("Execute", mock.Anything).Return(mockProducts, nil)

		controller := NewListProductsController(mockUseCase)

		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		w := httptest.NewRecorder()

		controller.Execute(w, req)

		resp := w.Result()
		defer resp.Body.Close()
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("should return 200 OK and empty list when no products", func(t *testing.T) {
		mockUseCase := &usecases_test.MockListProductsUseCase{}
		mockUseCase.On("Execute", mock.Anything).Return([]entities.Product{}, nil)

		controller := NewListProductsController(mockUseCase)

		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		w := httptest.NewRecorder()

		controller.Execute(w, req)

		resp := w.Result()
		defer resp.Body.Close()
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("should call HandleErrors and return error status when use case returns error", func(t *testing.T) {
		mockUseCase := &usecases_test.MockListProductsUseCase{}
		mockUseCase.On("Execute", mock.Anything).Return([]entities.Product{}, errors.New("unexpected error"))

		controller := NewListProductsController(mockUseCase)

		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		w := httptest.NewRecorder()

		controller.Execute(w, req)

		resp := w.Result()
		defer resp.Body.Close()
		assert.NotEqual(t, http.StatusOK, resp.StatusCode)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("should pass correct request to use case", func(t *testing.T) {
		mockUseCase := &usecases_test.MockListProductsUseCase{}
		mockUseCase.On("Execute", mock.AnythingOfType("usecases.ListProductsRequest")).Return([]entities.Product{}, nil)

		controller := NewListProductsController(mockUseCase)

		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		w := httptest.NewRecorder()

		controller.Execute(w, req)

		mockUseCase.AssertCalled(t, "Execute", mock.AnythingOfType("usecases.ListProductsRequest"))
	})
}
