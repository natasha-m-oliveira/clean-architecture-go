package usecases

import (
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/errors"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/repositories"
)

type (
	GetCartByIdUseCase interface {
		Execute(request GetCartByIdRequest) (*entities.Cart, error)
	}

	GetCartByIdRequest struct {
		Id string
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

func (uc getCartByIdUseCase) Execute(request GetCartByIdRequest) (*entities.Cart, error) {
	cart, err := uc.cartsRepository.FindById(request.Id)
	if err != nil {
		return nil, err
	}

	if cart == nil {
		return nil, errors.NewCartNotFound()
	}

	return cart, nil
}
