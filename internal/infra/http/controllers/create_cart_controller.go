package controllers

import (
	"net/http"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/usecases"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/dtos/input"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/handlers"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/mappers"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/render"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/utils"
)

type CreateCartController struct {
	createCartUseCase usecases.CreateCartUseCase
	cartMapper     mappers.HttpCartMapper
}

func NewCreateCartController(createCartUseCase usecases.CreateCartUseCase) CreateCartController {
	return CreateCartController{
		createCartUseCase: createCartUseCase,
		cartMapper:     mappers.NewHttpCartMapper(),
	}
}

func (c CreateCartController) Execute(w http.ResponseWriter, r *http.Request) {
	createCartInput, err := utils.DecodeBody(r.Body, input.CreateCartInput{})
	if err != nil {
		render.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	if err := createCartInput.Validate(); err != nil {
		render.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	createCartResponse, err := c.createCartUseCase.Execute(usecases.CreateCartRequest{
		Items: []struct {
			ProductId string
			Quantity  int
		}(createCartInput.Items),
	})
	if err != nil {
		handlers.HandleErrors(w, err)
		return
	}

	output := c.cartMapper.ToHttp(createCartResponse.Cart)

	render.NewSuccess(output, http.StatusCreated).Send(w)
}
