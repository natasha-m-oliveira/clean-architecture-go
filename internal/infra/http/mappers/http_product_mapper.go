package mappers

import (
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/dtos/output"
)

type HttpProductMapper struct{}

func NewHttpProductMapper() HttpProductMapper {
	return HttpProductMapper{}
}

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
