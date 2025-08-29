package usecases

import (
	"github.com/natasha-m-oliveira/clean-architecture-go/core/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/core/errors"
	"github.com/natasha-m-oliveira/clean-architecture-go/core/repositories"
)

type (
	CreateCartUseCase interface {
		Execute(request CreateCartRequest) (*entities.Cart, error)
	}

	CreateCartItemRequest struct {
		ProductId string
		Quantity  int
	}

	CreateCartRequest struct {
		Items []CreateCartItemRequest
	}
	createCartUseCase struct {
		cartsRepository    repositories.CartsRepository
		productsRepository repositories.ProductsRepository
	}
)

func NewCreateCartUseCase(
	cartsRepository repositories.CartsRepository,
	productsRepository repositories.ProductsRepository,
) CreateCartUseCase {
	return &createCartUseCase{
		cartsRepository:    cartsRepository,
		productsRepository: productsRepository,
	}
}

func (uc createCartUseCase) Execute(request CreateCartRequest) (*entities.Cart, error) {
	cart := entities.NewCart(entities.CartStatusPending, make([]entities.CartItem, len(request.Items)))

	for index, item := range request.Items {
		product, err := uc.productsRepository.FindById(item.ProductId)
		if err != nil {
			return nil, err
		}

		if product == nil {
			return nil, errors.NewProductNotFound()
		}

		cartItem := entities.NewCartItem(cart.Id, item.ProductId, item.Quantity, entities.CartItemOptions{
			Product: *product,
		})

		cart.Items[index] = *cartItem
	}

	err := uc.cartsRepository.Create(cart)
	if err != nil {
		return nil, err
	}

	return cart, nil
}
