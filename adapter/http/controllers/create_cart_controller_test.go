package controllers

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/natasha-m-oliveira/clean-architecture-go/adapter/http/input"
	"github.com/natasha-m-oliveira/clean-architecture-go/core/entities"
	usecases_test "github.com/natasha-m-oliveira/clean-architecture-go/test/usecases"
	"github.com/natasha-m-oliveira/clean-architecture-go/test/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCartController(t *testing.T) {
	t.Run("should return 201 Created when cart is created successfully", func(t *testing.T) {
		mockUseCase := &usecases_test.MockCreateCartUseCase{}
		mockUseCase.On("Execute", mock.Anything).Return(&entities.Cart{
			Id: "cart1",
			Items: []entities.CartItem{
				{ProductId: "069ea40a-392d-5719-8419-c04a93a773ed", Quantity: 2},
			},
		}, nil)

		controller := NewCreateCartController(mockUseCase)

		inputBody := input.CreateCartInput{
			Items: []input.CreateCartItemInput{
				{ProductId: "069ea40a-392d-5719-8419-c04a93a773ed", Quantity: 2},
			},
		}

		req := httptest.NewRequest(http.MethodPost, "/cart", utils.MakeRequestBody(t, inputBody))
		w := httptest.NewRecorder()

		controller.Execute(w, req)

		resp := w.Result()
		defer resp.Body.Close()
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("should return 400 Bad Request when request body is invalid JSON", func(t *testing.T) {
		mockUseCase := &usecases_test.MockCreateCartUseCase{}

		controller := NewCreateCartController(mockUseCase)

		req := httptest.NewRequest(http.MethodPost, "/cart", bytes.NewReader([]byte("{invalid json")))
		w := httptest.NewRecorder()

		controller.Execute(w, req)

		resp := w.Result()
		defer resp.Body.Close()
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		mockUseCase.AssertNotCalled(t, "Execute", mock.Anything)
	})

	t.Run("should return 400 Bad Request when input validation fails (empty items)", func(t *testing.T) {
		mockUseCase := &usecases_test.MockCreateCartUseCase{}

		controller := NewCreateCartController(mockUseCase)

		inputBody := input.CreateCartInput{
			Items: []input.CreateCartItemInput{},
		}
		req := httptest.NewRequest(http.MethodPost, "/cart", utils.MakeRequestBody(t, inputBody))
		w := httptest.NewRecorder()

		controller.Execute(w, req)

		resp := w.Result()
		defer resp.Body.Close()
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		mockUseCase.AssertNotCalled(t, "Execute", mock.Anything)
	})

	t.Run("should return error status when use case returns an error", func(t *testing.T) {
		mockUseCase := &usecases_test.MockCreateCartUseCase{}
		mockUseCase.On("Execute", mock.Anything).Return(nil, errors.New("use case error"))

		controller := NewCreateCartController(mockUseCase)

		inputBody := input.CreateCartInput{
			Items: []input.CreateCartItemInput{
				{ProductId: "0b750668-efce-5135-91e2-10803e2e77aa", Quantity: 2},
			},
		}

		req := httptest.NewRequest(http.MethodPost, "/cart", utils.MakeRequestBody(t, inputBody))
		w := httptest.NewRecorder()

		controller.Execute(w, req)

		resp := w.Result()
		defer resp.Body.Close()
		assert.NotEqual(t, http.StatusCreated, resp.StatusCode)
		mockUseCase.AssertExpectations(t)
	})
}
