package controllers

import (
	"net/http"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/adapter/dtos/input"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/adapter/handler"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/adapter/response"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/adapter/response/mappers"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/adapter/utils"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/usecases"
)

type CreateProductController struct {
	createProductUseCase usecases.CreateProductUseCase
}

func NewCreateProductController(
	createCartUseCase usecases.CreateProductUseCase,
) CreateProductController {
	return CreateProductController{
		createProductUseCase: createCartUseCase,
	}
}

func (c CreateProductController) Execute(w http.ResponseWriter, r *http.Request) {
	createProductInput, err := utils.DecodeBody(r.Body, input.CreateProductInput{})
	if err != nil {
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	if err := createProductInput.Validate(); err != nil {
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	createProductResponse, err := c.createProductUseCase.Execute(usecases.CreateProductRequest{
		Name:        createProductInput.Name,
		Description: createProductInput.Description,
		Price:       createProductInput.Price,
		Discount:    createProductInput.Discount,
	})
	if err != nil {
		handler.HandleErrors(w, err)
		return
	}

	output := (mappers.HttpProductMapper{}).ToHttp(createProductResponse.Product)

	response.NewSuccess(output, http.StatusCreated).Send(w)
}
