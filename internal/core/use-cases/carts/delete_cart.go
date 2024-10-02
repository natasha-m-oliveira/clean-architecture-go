package carts

import "github.com/natasha-m-oliveira/clean-architecture-go/internal/core/repositories"

type DeleteCartRequest struct {
	Id string
}

type DeleteCartUseCase struct {
	cartsRepository repositories.CartsRepository
}

func NewDeleteCartUseCase(cartsRepository repositories.CartsRepository) *DeleteCartUseCase {
	return &DeleteCartUseCase{
		cartsRepository: cartsRepository,
	}
}

func (uc *DeleteCartUseCase) Execute(request DeleteCartRequest) error {
	_, err := uc.cartsRepository.FindById(request.Id)
	if err != nil {
		return err
	}

	err = uc.cartsRepository.DeleteById(request.Id)

	return err
}
