package usecases

import (
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/errors"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/repositories"
)

type (
	UpdateCartItemsUseCase interface {
		Execute(request UpdateCartItemsRequest) (*entities.Cart, error)
	}

	UpdateCartItemsRequest struct {
		Id    string
		Items []struct {
			ProductId string
			Quantity  int
		}
	}

	updateCartItemsUseCase struct {
		cartsRepository    repositories.CartsRepository
		productsRepository repositories.ProductsRepository
	}
)

func NewUpdateCartItemsUseCase(cartsRepository repositories.CartsRepository, productsRepository repositories.ProductsRepository) UpdateCartItemsUseCase {
	return &updateCartItemsUseCase{
		cartsRepository:    cartsRepository,
		productsRepository: productsRepository,
	}
}

func (uc updateCartItemsUseCase) Execute(request UpdateCartItemsRequest) (*entities.Cart, error) {
	cart, err := uc.cartsRepository.FindById(request.Id)
	if err != nil {
		return nil, err
	}

	if cart == nil {
		return nil, errors.NewCartNotFound()
	}

	cart.Items = make([]entities.CartItem, len(request.Items))

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

	err = uc.cartsRepository.Save(cart)
	if err != nil {
		return nil, err
	}

	return cart, nil
}
