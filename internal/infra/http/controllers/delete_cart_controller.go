package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/usecases"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/errors"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/render"
)

type DeleteCartController struct {
	deleteCartUseCase usecases.DeleteCartUseCase
}

func NewDeleteCartController(
	deleteCartUseCase usecases.DeleteCartUseCase,
) DeleteCartController {
	return DeleteCartController{
		deleteCartUseCase: deleteCartUseCase,
	}
}

func (c DeleteCartController) Execute(w http.ResponseWriter, r *http.Request) {
	cartId := chi.URLParam(r, "id")

	err := c.deleteCartUseCase.Execute(usecases.DeleteCartRequest{
		Id: cartId,
	})

	if err != nil {
		errors.WriteError(w, err)
		return
	}

	render.NewResponse(nil, http.StatusNoContent).Send(w)
}
