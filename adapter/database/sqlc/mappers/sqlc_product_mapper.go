package mappers

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/natasha-m-oliveira/clean-architecture-go/adapter/utils"
	"github.com/natasha-m-oliveira/clean-architecture-go/core/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/sqlc/db"
)

type SqlcProductMapper struct{}

func (m SqlcProductMapper) ToSqlc(product entities.Product) db.Product {
	description := pgtype.Text{String: product.Description, Valid: product.Description != ""}
	image := pgtype.Text{String: product.Image, Valid: product.Image != ""}
	discount := pgtype.Int4{Int32: int32(product.Discount), Valid: product.Discount != 0}

	return db.Product{
		ID:          product.Id,
		Name:        product.Name,
		Description: description,
		Image:       image,
		Price:       product.Price,
		Discount:    discount,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func (m SqlcProductMapper) ToDomain(model db.Product) *entities.Product {
	name := model.Name
	price := int(model.Price)

	discount := utils.Ternary(model.Discount.Valid, int(model.Discount.Int32), 0)
	image := utils.Ternary(model.Image.Valid, model.Image.String, "")
	description := utils.Ternary(model.Description.Valid, model.Description.String, "")

	optional := entities.ProductOptions{
		Id:          model.ID,
		Description: description,
		Image:       image,
		Discount:    discount,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}

	product := entities.NewProduct(name, price, optional)

	return product
}
