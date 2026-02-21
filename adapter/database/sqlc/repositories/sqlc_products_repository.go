package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/natasha-m-oliveira/clean-architecture-go/adapter/database/sqlc/mappers"
	"github.com/natasha-m-oliveira/clean-architecture-go/adapter/utils"
	"github.com/natasha-m-oliveira/clean-architecture-go/core/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/core/repositories"
	"github.com/natasha-m-oliveira/clean-architecture-go/sqlc/db"
)

type SqlcProductsRepository struct {
	connection *pgx.Conn
	ctx        context.Context
	queries    *db.Queries
	mapper     mappers.SqlcProductMapper
}

func NewSqlcProductsRepository(connection *pgx.Conn, ctx context.Context) repositories.ProductsRepository {
	return &SqlcProductsRepository{
		connection: connection,
		ctx:        ctx,
		queries:    db.New(connection),
		mapper:     mappers.SqlcProductMapper{},
	}
}

func (r *SqlcProductsRepository) Create(product *entities.Product) error {
	params := db.CreateProductParams{}
	utils.Copy(&params, &product)

	return r.queries.CreateProduct(r.ctx, params)
}

func (r *SqlcProductsRepository) FindById(id string) (*entities.Product, error) {
	row, err := r.queries.GetProductById(r.ctx, id)
	if err != nil && err == pgx.ErrNoRows {
		return nil, nil
	}

	product := r.mapper.ToDomain(db.Product(row))

	return product, nil
}

func (r *SqlcProductsRepository) FindByName(name string) (*entities.Product, error) {
	row, err := r.queries.GetProductByName(r.ctx, name)
	if err != nil && err == pgx.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	product := r.mapper.ToDomain(db.Product(row))

	return product, nil
}

func (r *SqlcProductsRepository) List() ([]entities.Product, error) {
	rows, err := r.queries.ListProducts(r.ctx)
	if err != nil && err == pgx.ErrNoRows {
		return []entities.Product{}, nil
	}

	if err != nil {
		return nil, err
	}

	products := make([]entities.Product, 0, len(rows))

	for _, row := range rows {
		model := db.Product{}
		utils.Copy(&model, &row)

		product := r.mapper.ToDomain(model)
		products = append(products, *product)
	}

	return products, nil
}

func (r *SqlcProductsRepository) Save(product *entities.Product) error {
	params := db.UpdateProductParams{}
	utils.Copy(&params, &product)

	return r.queries.UpdateProduct(r.ctx, params)
}

func (r *SqlcProductsRepository) DeleteById(id string) error {
	return r.queries.DeleteProductById(r.ctx, id)
}
