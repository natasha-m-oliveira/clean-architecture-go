package mappers

import (
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/adapter/dtos/output"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/entities"
)

type HttpProductMapper struct{}

func (m HttpProductMapper) ToHttp(domain entities.Product) output.Product {
	return output.Product{
		Id:          domain.Id,
		Name:        domain.Name,
		Description: domain.Description,
		Image:       domain.Image,
		Price:       domain.Price,
		Discount:    domain.Discount,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
