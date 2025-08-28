package controllers

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/adapter/http/input"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/entities"
	usecases_test "github.com/natasha-m-oliveira/clean-architecture-go/test/usecases"
	"github.com/natasha-m-oliveira/clean-architecture-go/test/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateProductController(t *testing.T) {
	t.Run("should return 201 Created when product is created successfully", func(t *testing.T) {
		mockUseCase := &usecases_test.MockCreateProductUseCase{}
		mockUseCase.On("Execute", mock.Anything).Return(&entities.Product{
			Id:    "22475cef-4675-580f-8710-bbf8cd043fb5",
			Name:  "Test Product",
			Price: 105,
		}, nil)

		controller := NewCreateProductController(mockUseCase)

		inputBody := input.CreateProductInput{
			Name:        "Test Product",
			Description: "Test Description",
			Price:       105,
		}

		req := httptest.NewRequest(http.MethodPost, "/product", utils.MakeRequestBody(t, inputBody))
		w := httptest.NewRecorder()

		controller.Execute(w, req)

		resp := w.Result()
		defer resp.Body.Close()
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("should return 400 Bad Request when request body is invalid JSON", func(t *testing.T) {
		mockUseCase := &usecases_test.MockCreateProductUseCase{}
		controller := NewCreateProductController(mockUseCase)

		req := httptest.NewRequest(http.MethodPost, "/product", bytes.NewReader([]byte("{invalid json")))
		w := httptest.NewRecorder()

		controller.Execute(w, req)

		resp := w.Result()
		defer resp.Body.Close()
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		mockUseCase.AssertNotCalled(t, "Execute", mock.Anything)
	})

	t.Run("should return 400 Bad Request when input validation fails (empty name)", func(t *testing.T) {
		mockUseCase := &usecases_test.MockCreateProductUseCase{}
		controller := NewCreateProductController(mockUseCase)

		inputBody := input.CreateProductInput{
			Name:  "",
			Price: 105,
		}
		req := httptest.NewRequest(http.MethodPost, "/product", utils.MakeRequestBody(t, inputBody))
		w := httptest.NewRecorder()

		controller.Execute(w, req)

		resp := w.Result()
		defer resp.Body.Close()
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		mockUseCase.AssertNotCalled(t, "Execute", mock.Anything)
	})

	t.Run("should return 400 Bad Request when input validation fails (zero price)", func(t *testing.T) {
		mockUseCase := &usecases_test.MockCreateProductUseCase{}
		controller := NewCreateProductController(mockUseCase)

		inputBody := input.CreateProductInput{
			Name:        "Valid Name",
			Description: "Test Description",
			Price:       0,
		}
		req := httptest.NewRequest(http.MethodPost, "/product", utils.MakeRequestBody(t, inputBody))
		w := httptest.NewRecorder()

		controller.Execute(w, req)

		resp := w.Result()
		defer resp.Body.Close()
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		mockUseCase.AssertNotCalled(t, "Execute", mock.Anything)
	})

	t.Run("should return error status when use case returns an error", func(t *testing.T) {
		mockUseCase := &usecases_test.MockCreateProductUseCase{}
		mockUseCase.On("Execute", mock.Anything).Return(nil, errors.New("use case error"))

		controller := NewCreateProductController(mockUseCase)

		inputBody := input.CreateProductInput{
			Name:        "Test Product",
			Description: "Test Description",
			Price:       105,
		}

		req := httptest.NewRequest(http.MethodPost, "/product", utils.MakeRequestBody(t, inputBody))
		w := httptest.NewRecorder()

		controller.Execute(w, req)

		resp := w.Result()
		defer resp.Body.Close()
		assert.NotEqual(t, http.StatusCreated, resp.StatusCode)
		mockUseCase.AssertExpectations(t)
	})
}
