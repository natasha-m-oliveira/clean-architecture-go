package controllers

import (
	"net/http"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/usecases"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/dtos/output"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/errors"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/mappers"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/render"
)

type ListProductsController struct {
	listProductsUseCase usecases.ListProductsUseCase
	productMapper       mappers.HttpProductMapper
}

func NewListProductsController(listProductsUseCase usecases.ListProductsUseCase) ListProductsController {
	return ListProductsController{
		listProductsUseCase: listProductsUseCase,
		productMapper:       mappers.NewHttpProductMapper(),
	}
}

func (c ListProductsController) Execute(w http.ResponseWriter, r *http.Request) {
	createProductResponse, err := c.listProductsUseCase.Execute(usecases.ListProductsRequest{})
	if err != nil {
		errors.WriteError(w, err)
		return
	}

	output := make([]output.Product, len(createProductResponse.Products))

	for i, product := range createProductResponse.Products {
		output[i] = c.productMapper.ToHttp(product)
	}

	render.NewResponse(output, http.StatusOK).Send(w)
}
