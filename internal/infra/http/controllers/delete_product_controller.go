package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/usecases"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/handlers"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/render"
)

type DeleteProductController struct {
	deleteProductUseCase usecases.DeleteProductUseCase
}

func NewDeleteProductController(
	deleteProductUseCase usecases.DeleteProductUseCase,
) DeleteProductController {
	return DeleteProductController{
		deleteProductUseCase: deleteProductUseCase,
	}
}

func (c DeleteProductController) Execute(w http.ResponseWriter, r *http.Request) {
	productId := chi.URLParam(r, "id")

	err := c.deleteProductUseCase.Execute(usecases.DeleteProductRequest{
		Id: productId,
	})

	if err != nil {
		handlers.HandleErrors(w, err)
		return
	}

	render.NewSuccess(nil, http.StatusNoContent).Send(w)
}
