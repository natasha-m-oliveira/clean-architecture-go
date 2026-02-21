package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/usecases"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/handlers"
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
		handlers.HandleErrors(w, err)
		return
	}

	render.NewSuccess(nil, http.StatusNoContent).Send(w)
}
