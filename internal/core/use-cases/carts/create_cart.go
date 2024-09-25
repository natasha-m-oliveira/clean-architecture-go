package carts

import (
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/repositories"
)

type CreateCartRequest struct {
	Items []struct {
		ProductId string
		Quantity  int
	}
}

type CreateCartResponse struct {
	Cart entities.Cart
}

type CreateCartUseCase struct {
	cartsRepository    repositories.CartsRepository
	productsRepository repositories.ProductsRepository
}

func NewCreateCartUseCase(cartsRepository repositories.CartsRepository, productsRepository repositories.ProductsRepository) *CreateCartUseCase {
	return &CreateCartUseCase{
		cartsRepository:    cartsRepository,
		productsRepository: productsRepository,
	}
}

func (uc *CreateCartUseCase) Execute(request CreateCartRequest) (*CreateCartResponse, error) {
	var items []entities.CartItem

	for _, item := range request.Items {
		product, err := uc.productsRepository.FindById(item.ProductId)
		if err != nil {
			return nil, err
		}

		cartItem := entities.NewCartItem("", item.ProductId, item.Quantity, entities.CartItemOptions{
			Product: *product,
		})

		items = append(items, *cartItem)
	}

	cart := entities.NewCart(entities.CartStatusPending, items)

	err := uc.cartsRepository.Create(cart)
	if err != nil {
		return nil, err
	}

	return &CreateCartResponse{
		Cart: *cart,
	}, nil
}
