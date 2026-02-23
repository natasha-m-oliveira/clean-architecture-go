package repositories

import (
	"context"
	"errors"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/entities"
	_errors "github.com/natasha-m-oliveira/clean-architecture-go/internal/app/errors"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/database/prisma/repositories/mappers"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/utils"
	"github.com/natasha-m-oliveira/clean-architecture-go/prisma/db"
)

type PrismaProductsRepository struct {
	ctx           context.Context
	client        *db.PrismaClient
	productMapper mappers.PrismaProductMapper
}

func NewPrismaProductsRepository(ctx context.Context, client *db.PrismaClient) *PrismaProductsRepository {
	return &PrismaProductsRepository{
		ctx:           ctx,
		client:        client,
		productMapper: mappers.NewPrismaProductMapper(),
	}
}

func (r *PrismaProductsRepository) Create(product *entities.Product) error {
	name, price, optional := r.productMapper.ToPrisma(*product)

	_, err := r.client.Product.CreateOne(name, price, optional...).Exec(r.ctx)
	if err != nil {
		return _errors.NewInternalServerError()
	}

	return nil
}

func (r *PrismaProductsRepository) FindById(id string) (*entities.Product, error) {
	model, err := r.client.Product.FindUnique(db.Product.ID.Equals(id)).Exec(r.ctx)
	if err != nil && errors.Is(err, db.ErrNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, _errors.NewInternalServerError()
	}

	product := r.productMapper.ToDomain(*model)

	return &product, err
}

func (r *PrismaProductsRepository) FindByName(name string) (*entities.Product, error) {
	model, err := r.client.Product.FindUnique(db.Product.Name.Equals(name)).Exec(r.ctx)
	if err != nil && errors.Is(err, db.ErrNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, _errors.NewInternalServerError()
	}

	product := r.productMapper.ToDomain(*model)

	return &product, err
}

func (r *PrismaProductsRepository) FindAll() ([]entities.Product, error) {
	models, err := r.client.Product.FindMany().Exec(r.ctx)
	if err != nil {
		return []entities.Product{}, _errors.NewInternalServerError()
	}

	products := utils.Map(models, r.productMapper.ToDomain)

	return products, nil
}

func (r *PrismaProductsRepository) Save(product *entities.Product) error {
	name, price, optional := r.productMapper.ToPrisma(*product)

	optional = append(optional, name.(db.ProductSetParam), price.(db.ProductSetParam))

	_, err := r.client.Product.FindUnique(db.Product.ID.Equals(product.Id)).Update(optional...).Exec(r.ctx)
	if err != nil {
		return _errors.NewInternalServerError()
	}
	return nil
}

func (r *PrismaProductsRepository) DeleteById(id string) error {
	_, err := r.client.Product.FindUnique(db.Product.ID.Equals(id)).Delete().Exec(r.ctx)
	if err != nil {
		return _errors.NewInternalServerError()
	}
	return nil
}
