package usecases

import (
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/repositories"
)

type (
	CreateCartUseCase interface {
		Execute(request CreateCartRequest) (*CreateCartResponse, error)
	}

	CreateCartRequest struct {
		Items []struct {
			ProductId string
			Quantity  int
		}
	}

	CreateCartResponse struct {
		Cart entities.Cart
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

func (uc createCartUseCase) Execute(request CreateCartRequest) (*CreateCartResponse, error) {
	cart := entities.NewCart(entities.CartStatusPending, make([]entities.CartItem, len(request.Items)))

	for index, item := range request.Items {
		product, err := uc.productsRepository.FindById(item.ProductId)
		if err != nil {
			return nil, err
		}

		cartItem := entities.NewCartItem("", item.ProductId, item.Quantity, entities.CartItemOptions{
			Product: *product,
		})

		cart.Items[index] = *cartItem
	}

	err := uc.cartsRepository.Create(cart)
	if err != nil {
		return nil, err
	}

	return &CreateCartResponse{
		Cart: *cart,
	}, nil
}
