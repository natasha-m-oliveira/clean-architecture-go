package controllers

import (
	"net/http"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/usecases"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/dtos/input"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/handlers"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/mappers"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/render"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/utils"
)

type CreateProductController struct {
	createProductUseCase usecases.CreateProductUseCase
	productMapper        mappers.HttpProductMapper
}

func NewCreateProductController(createProductUseCase usecases.CreateProductUseCase) CreateProductController {
	return CreateProductController{
		createProductUseCase: createProductUseCase,
		productMapper:        mappers.NewHttpProductMapper(),
	}
}

func (c CreateProductController) Execute(w http.ResponseWriter, r *http.Request) {
	createProductInput, err := utils.DecodeBody(r.Body, input.CreateProductInput{})
	if err != nil {
		render.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	if err := createProductInput.Validate(); err != nil {
		render.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	createProductResponse, err := c.createProductUseCase.Execute(usecases.CreateProductRequest{
		Name:        createProductInput.Name,
		Description: createProductInput.Description,
		Price:       createProductInput.Price,
		Discount:    createProductInput.Discount,
	})
	if err != nil {
		handlers.HandleErrors(w, err)
		return
	}

	output := c.productMapper.ToHttp(createProductResponse.Product)

	render.NewSuccess(output, http.StatusCreated).Send(w)
}
