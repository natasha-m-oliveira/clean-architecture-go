package repositories

import (
	"context"
	"errors"

	"github.com/natasha-m-oliveira/clean-architecture-go/adapter/database/prisma/mappers"
	"github.com/natasha-m-oliveira/clean-architecture-go/adapter/utils"
	"github.com/natasha-m-oliveira/clean-architecture-go/core/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/prisma/db"
)

type PrismaProductsRepository struct {
	client *db.PrismaClient
	ctx    context.Context
}

func NewPrismaProductsRepository(client *db.PrismaClient, ctx context.Context) *PrismaProductsRepository {
	return &PrismaProductsRepository{
		client: client,
		ctx:    ctx,
	}
}

func (r *PrismaProductsRepository) Create(product *entities.Product) error {
	name, price, optional := (mappers.PrismaProductMapper{}).ToPrisma(*product)

	_, err := r.client.Product.CreateOne(name, price, optional...).Exec(r.ctx)

	return err
}

func (r *PrismaProductsRepository) FindById(id string) (*entities.Product, error) {
	model, err := r.client.Product.FindUnique(db.Product.ID.Equals(id)).Exec(r.ctx)
	if err != nil && errors.Is(err, db.ErrNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	product := (mappers.PrismaProductMapper{}).ToDomain(*model)

	return &product, err
}

func (r *PrismaProductsRepository) FindByName(name string) (*entities.Product, error) {
	model, err := r.client.Product.FindUnique(db.Product.Name.Equals(name)).Exec(r.ctx)
	if err != nil && errors.Is(err, db.ErrNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	product := (mappers.PrismaProductMapper{}).ToDomain(*model)

	return &product, err
}

func (r *PrismaProductsRepository) List() ([]entities.Product, error) {
	models, err := r.client.Product.FindMany().Exec(r.ctx)
	if err != nil {
		return []entities.Product{}, err
	}

	products := utils.MapToDomain(models, (mappers.PrismaProductMapper{}).ToDomain)

	return products, nil
}

func (r *PrismaProductsRepository) Save(product *entities.Product) error {
	name, price, optional := (mappers.PrismaProductMapper{}).ToPrisma(*product)

	optional = append(optional, name.(db.ProductSetParam), price.(db.ProductSetParam))

	_, err := r.client.Product.FindUnique(db.Product.ID.Equals(product.Id)).Update(optional...).Exec(r.ctx)

	return err
}

func (r *PrismaProductsRepository) DeleteById(id string) error {
	_, err := r.client.Product.FindUnique(db.Product.ID.Equals(id)).Delete().Exec(r.ctx)

	return err
}
