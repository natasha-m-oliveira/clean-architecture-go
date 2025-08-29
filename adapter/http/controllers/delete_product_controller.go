package controllers

import (
	"net/http"

	"github.com/natasha-m-oliveira/clean-architecture-go/adapter/http/handlers"
	"github.com/natasha-m-oliveira/clean-architecture-go/adapter/http/response"
	"github.com/natasha-m-oliveira/clean-architecture-go/core/usecases"
)

type DeleteProductController struct {
	deleteProductUseCase usecases.DeleteProductUseCase
}

func NewDeleteProductController(
	deleteCartUseCase usecases.DeleteProductUseCase,
) DeleteProductController {
	return DeleteProductController{
		deleteProductUseCase: deleteCartUseCase,
	}
}

func (c DeleteProductController) Execute(w http.ResponseWriter, r *http.Request) {
	req := usecases.DeleteProductRequest{
		Id: r.PathValue("id"),
	}

	err := c.deleteProductUseCase.Execute(req)
	if err != nil {
		handlers.HandleErrors(w, err)
		return
	}

	response.NewSuccess(nil, http.StatusNoContent).Send(w)
}
