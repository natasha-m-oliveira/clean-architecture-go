package controllers

import (
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/natasha-m-oliveira/clean-architecture-go/adapter/http/handlers"
	"github.com/natasha-m-oliveira/clean-architecture-go/adapter/http/output"
	"github.com/natasha-m-oliveira/clean-architecture-go/adapter/http/response"
	"github.com/natasha-m-oliveira/clean-architecture-go/core/usecases"
)

type ListProductsController struct {
	listProductsUseCase usecases.ListProductsUseCase
}

func NewListProductsController(
	createCartUseCase usecases.ListProductsUseCase,
) ListProductsController {
	return ListProductsController{
		listProductsUseCase: createCartUseCase,
	}
}

func (c ListProductsController) Execute(w http.ResponseWriter, r *http.Request) {
	req := usecases.ListProductsRequest{}

	products, err := c.listProductsUseCase.Execute(req)
	if err != nil {
		handlers.HandleErrors(w, err)
		return
	}

	var res []output.Product
	copier.Copy(&res, &products)

	response.NewSuccess(res, http.StatusOK).Send(w)
}
