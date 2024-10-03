package mappers

import (
	"testing"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/prisma/db"
)

func TestPrismaProductMapper_ToPrisma(t *testing.T) {
	mapper := PrismaProductMapper{}

	product := entities.NewProduct("Test Product", 999)

	name, price, optional := mapper.ToPrisma(*product)

	expectedName := db.ProductWithPrismaNameSetParam("Test Product")
	if name != expectedName {
		t.Errorf("Expected name to be %s, but got %s", expectedName, name)
	}

	expectedPrice := db.ProductWithPrismaPriceSetParam(9.99)
	if price != expectedPrice {
		t.Errorf("Expected price to be %f, but got %f", expectedPrice, price)
	}

	if len(optional) != 0 {
		t.Errorf("Expected optional to be empty, but got %d elements", len(optional))
	}
}

func TestPrismaProductMapper_ToDomain(t *testing.T) {
	mapper := PrismaProductMapper{}

	model := db.ProductModel{
		ID:    "b3cd9263-5d04-5bb3-ae9e-691b2240465d",
		Name:  "Test Product",
		Price: 999,
	}

	product := mapper.ToDomain(model)

	expectedProduct := entities.Product{
		Id:    "b3cd9263-5d04-5bb3-ae9e-691b2240465d",
		Name:  "Test Product",
		Price: 999,
	}

	if product != expectedProduct {
		t.Errorf("Expected product to be %+v, but got %+v", expectedProduct, product)
	}
}
