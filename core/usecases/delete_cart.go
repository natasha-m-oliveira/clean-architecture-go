package usecases

import (
	"github.com/natasha-m-oliveira/clean-architecture-go/core/errors"
	"github.com/natasha-m-oliveira/clean-architecture-go/core/repositories"
)

type (
	DeleteCartUseCase interface {
		Execute(request DeleteCartRequest) error
	}

	DeleteCartRequest struct {
		Id string
	}

	deleteCartUseCase struct {
		cartsRepository repositories.CartsRepository
	}
)

func NewDeleteCartUseCase(cartsRepository repositories.CartsRepository) DeleteCartUseCase {
	return &deleteCartUseCase{
		cartsRepository: cartsRepository,
	}
}

func (uc deleteCartUseCase) Execute(request DeleteCartRequest) error {
	cart, err := uc.cartsRepository.FindById(request.Id)
	if err != nil {
		return err
	}

	if cart == nil {
		return errors.NewCartNotFound()
	}

	return uc.cartsRepository.DeleteById(request.Id)
}
