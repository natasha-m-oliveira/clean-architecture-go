package controllers

import (
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/adapter/http/handlers"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/adapter/http/input"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/adapter/http/output"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/adapter/http/response"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/adapter/utils"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/usecases"
)

type CreateCartController struct {
	createCartUseCase usecases.CreateCartUseCase
}

func NewCreateCartController(
	createCartUseCase usecases.CreateCartUseCase,
) CreateCartController {
	return CreateCartController{
		createCartUseCase: createCartUseCase,
	}
}

func (c CreateCartController) Execute(w http.ResponseWriter, r *http.Request) {
	input, err := utils.DecodeBody(r.Body, input.CreateCartInput{})
	if err != nil {
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	if err := input.Validate(); err != nil {
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	var items []usecases.CreateCartItemRequest
	copier.Copy(&items, &input.Items)

	req := usecases.CreateCartRequest{
		Items: items,
	}

	cart, err := c.createCartUseCase.Execute(req)
	if err != nil {
		handlers.HandleErrors(w, err)
		return
	}

	var res output.Cart
	copier.Copy(&res, &cart)

	response.NewSuccess(res, http.StatusCreated).Send(w)
}
