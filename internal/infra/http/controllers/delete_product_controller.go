package controllers

import (
	"net/http"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/usecases"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/errors"
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
	productId := r.PathValue("id")

	err := c.deleteProductUseCase.Execute(usecases.DeleteProductRequest{
		Id: productId,
	})

	if err != nil {
		errors.WriteError(w, err)
		return
	}

	render.NewResponse(nil, http.StatusNoContent).Send(w)
}
