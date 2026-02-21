package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/usecases"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/handlers"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/mappers"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/render"
)

type GetProductByIdController struct {
	getProductByIdUseCase usecases.GetProductByIdUseCase
	productMapper         mappers.HttpProductMapper
}

func NewGetProductByIdController(getProductByIdUseCase usecases.GetProductByIdUseCase) GetProductByIdController {
	return GetProductByIdController{
		getProductByIdUseCase: getProductByIdUseCase,
		productMapper:         mappers.NewHttpProductMapper(),
	}
}

func (c GetProductByIdController) Execute(w http.ResponseWriter, r *http.Request) {
	productId := chi.URLParam(r, "id")

	createProductResponse, err := c.getProductByIdUseCase.Execute(usecases.GetProductByIdRequest{
		Id: productId,
	})

	if err != nil {
		handlers.HandleErrors(w, err)
		return
	}

	output := c.productMapper.ToHttp(createProductResponse.Product)

	render.NewSuccess(output, http.StatusOK).Send(w)
}
