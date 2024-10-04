package usecases

import (
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/repositories"
)

type (
	GetCartByIdUseCase interface {
		Execute(request GetCartByIdRequest) (*GetCartByIdResponse, error)
	}

	GetCartByIdRequest struct {
		Id string
	}

	GetCartByIdResponse struct {
		Cart entities.Cart
	}

	getCartByIdUseCase struct {
		cartsRepository repositories.CartsRepository
	}
)

func NewGetCartByIdUseCase(cartsRepository repositories.CartsRepository) GetCartByIdUseCase {
	return &getCartByIdUseCase{
		cartsRepository: cartsRepository,
	}
}

func (uc getCartByIdUseCase) Execute(request GetCartByIdRequest) (*GetCartByIdResponse, error) {
	cart, err := uc.cartsRepository.FindById(request.Id)
	if err != nil {
		return nil, err
	}

	return &GetCartByIdResponse{
		Cart: *cart,
	}, nil
}
