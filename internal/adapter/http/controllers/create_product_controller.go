package controllers

import (
	"net/http"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/adapter/http/handlers"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/adapter/http/input"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/adapter/http/output"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/adapter/http/response"
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
	input, err := utils.DecodeBody(r.Body, input.CreateProductInput{})
	if err != nil {
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	if err := input.Validate(); err != nil {
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	req := usecases.CreateProductRequest(*input)

	product, err := c.createProductUseCase.Execute(req)
	if err != nil {
		handlers.HandleErrors(w, err)
		return
	}

	res := output.Product(*product)

	response.NewSuccess(res, http.StatusCreated).Send(w)
}
