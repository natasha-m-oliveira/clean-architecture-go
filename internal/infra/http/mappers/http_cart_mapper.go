package mappers

import (
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/dtos/output"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/utils"
)

type HttpCartMapper struct{}

func NewHttpCartMapper() HttpCartMapper {
	return HttpCartMapper{}
}

func (m HttpCartMapper) ToHttp(cart entities.Cart) output.Cart {
	return output.Cart{
		Id:     cart.Id,
		Status: cart.Status.String(),
		Items: utils.Map(cart.Items, func(item entities.CartItem) output.CartItem {
			return output.CartItem{
				Id:        item.Id,
				ProductId: item.ProductId,
				Quantity:  item.Quantity,
				Product: output.Product{
					Id:          item.Product.Id,
					Name:        item.Product.Name,
					Description: item.Product.Description,
					Price:       item.Product.Price,
				},
			}
		}),
	}
}
