package controllers

import (
	"net/http"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/adapter/http/handlers"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/adapter/http/output"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/adapter/http/response"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/usecases"
)

type GetProductByIdController struct {
	getProductByIdUseCase usecases.GetProductByIdUseCase
}

func NewGetProductByIdController(
	createCartUseCase usecases.GetProductByIdUseCase,
) GetProductByIdController {
	return GetProductByIdController{
		getProductByIdUseCase: createCartUseCase,
	}
}

func (c GetProductByIdController) Execute(w http.ResponseWriter, r *http.Request) {
	req := usecases.GetProductByIdRequest{
		Id: r.PathValue("id"),
	}

	product, err := c.getProductByIdUseCase.Execute(req)
	if err != nil {
		handlers.HandleErrors(w, err)
		return
	}

	res := output.Product(*product)

	response.NewSuccess(res, http.StatusOK).Send(w)
}
