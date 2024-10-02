package carts

import (
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/repositories"
)

type GetCartByIdRequest struct {
	Id string
}

type GetCartByIdResponse struct {
	Cart entities.Cart
}

type GetCartByIdUseCase struct {
	cartsRepository repositories.CartsRepository
}

func NewGetCartByIdUseCase(cartsRepository repositories.CartsRepository) *GetCartByIdUseCase {
	return &GetCartByIdUseCase{
		cartsRepository: cartsRepository,
	}
}

func (uc *GetCartByIdUseCase) Execute(request GetCartByIdRequest) (*GetCartByIdResponse, error) {
	cart, err := uc.cartsRepository.FindById(request.Id)
	if err != nil {
		return nil, err
	}

	return &GetCartByIdResponse{
		Cart: *cart,
	}, nil
}
