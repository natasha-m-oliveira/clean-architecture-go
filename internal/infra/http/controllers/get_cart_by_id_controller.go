package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/usecases"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/handlers"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/mappers"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/render"
)

type GetCartByIdController struct {
	getCartByIdUseCase usecases.GetCartByIdUseCase
	cartMapper         mappers.HttpCartMapper
}

func NewGetCartByIdController(getCartByIdUseCase usecases.GetCartByIdUseCase) GetCartByIdController {
	return GetCartByIdController{
		getCartByIdUseCase: getCartByIdUseCase,
		cartMapper:         mappers.NewHttpCartMapper(),
	}
}

func (c GetCartByIdController) Execute(w http.ResponseWriter, r *http.Request) {
	cartId := chi.URLParam(r, "id")

	createCartResponse, err := c.getCartByIdUseCase.Execute(usecases.GetCartByIdRequest{
		Id: cartId,
	})

	if err != nil {
		handlers.HandleErrors(w, err)
		return
	}

	output := c.cartMapper.ToHttp(createCartResponse.Cart)

	render.NewSuccess(output, http.StatusOK).Send(w)
}
