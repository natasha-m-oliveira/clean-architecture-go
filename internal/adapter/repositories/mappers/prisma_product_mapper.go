package mappers

import (
	"reflect"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/prisma/db"
)

type PrismaProductMapper struct{}

func (m PrismaProductMapper) ToPrisma(product entities.Product) (name db.ProductWithPrismaNameSetParam, price db.ProductWithPrismaPriceSetParam, optional []db.ProductSetParam) {
	name = db.Product.Name.Set(product.Name)
	price = db.Product.Price.Set(product.Price)

	totalFields := reflect.TypeOf(entities.Product{}).NumField()
	optional = make([]db.ProductSetParam, 0, totalFields-2)

	optional = append(optional,
		db.Product.ID.SetIfPresent(&product.Id),
		db.Product.Description.SetIfPresent(&product.Description),
		db.Product.Image.SetIfPresent(&product.Image),
		db.Product.Discount.SetIfPresent(&product.Discount),
		db.Product.CreatedAt.Set(product.CreatedAt),
		db.Product.UpdatedAt.Set(product.UpdatedAt))

	return name, price, optional
}

func (m PrismaProductMapper) ToDomain(model db.ProductModel) entities.Product {
	name := model.Name
	price := model.Price

	description, _ := model.Description()
	image, _ := model.Image()
	discount, _ := model.Discount()

	optional := entities.ProductOptions{
		Id:          model.ID,
		Description: description,
		Image:       image,
		Discount:    discount,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}

	product := entities.NewProduct(name, price, optional)

	return *product
}
