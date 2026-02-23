package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/usecases"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/dtos/input"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/errors"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/mappers"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/render"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/utils"
)

type UpdateCartItemsController struct {
	updateCartItemsUseCase usecases.UpdateCartItemsUseCase
	cartMapper             mappers.HttpCartMapper
}

func NewUpdateCartItemsController(
	updateCartItemsUseCase usecases.UpdateCartItemsUseCase,
) UpdateCartItemsController {
	return UpdateCartItemsController{
		updateCartItemsUseCase: updateCartItemsUseCase,
		cartMapper:             mappers.NewHttpCartMapper(),
	}
}

func (c UpdateCartItemsController) Execute(w http.ResponseWriter, r *http.Request) {
	cartId := chi.URLParam(r, "id")

	updateCartInput, err := utils.DecodeBody(r.Body, input.UpdateCartItemsInput{})
	if err != nil {
		errors.WriteError(w, err)
		return
	}

	if err := updateCartInput.Validate(); err != nil {
		errors.WriteError(w, err)
		return
	}

	updateCartItemsResponse, err := c.updateCartItemsUseCase.Execute(usecases.UpdateCartItemsRequest{
		Id: cartId,
		Items: []struct {
			ProductId string
			Quantity  int
		}(updateCartInput.Items),
	})

	if err != nil {
		errors.WriteError(w, err)
		return
	}

	output := c.cartMapper.ToHttp(updateCartItemsResponse.Cart)

	render.NewResponse(output, http.StatusOK).Send(w)
}
