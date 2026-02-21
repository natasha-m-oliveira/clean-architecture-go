package repositories

import (
	"context"
	"encoding/json"

	"github.com/jackc/pgx/v5"
	"github.com/natasha-m-oliveira/clean-architecture-go/adapter/database/sqlc/mappers"
	"github.com/natasha-m-oliveira/clean-architecture-go/adapter/utils"
	"github.com/natasha-m-oliveira/clean-architecture-go/core/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/core/repositories"
	"github.com/natasha-m-oliveira/clean-architecture-go/sqlc/db"
)

type SqlcCartsRepository struct {
	connection *pgx.Conn
	ctx        context.Context
	queries    *db.Queries
	mapper     mappers.SqlcCartMapper
}

func NewSqlcCartsRepository(connection *pgx.Conn, ctx context.Context) repositories.CartsRepository {
	return &SqlcCartsRepository{
		connection: connection,
		ctx:        ctx,
		queries:    db.New(connection),
		mapper:     mappers.SqlcCartMapper{},
	}
}

func (r *SqlcCartsRepository) Create(cart *entities.Cart) error {
	model := r.mapper.ToSqlc(*cart)

	params := db.CreateCartParams{}
	utils.Copy(&params, &model)

	return r.queries.CreateCart(r.ctx, params)
}

func (r *SqlcCartsRepository) FindById(id string) (*entities.Cart, error) {
	row, err := r.queries.GetCartByIdWithItems(r.ctx, id)
	if err != nil && err == pgx.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	var CardItems []entities.CartItem
	if err := json.Unmarshal(row.Items, &CardItems); err != nil {
		return nil, err
	}

	model := db.Cart{}
	utils.Copy(&model, &row)

	cart := r.mapper.ToDomain(model)
	cart.Items = CardItems

	return cart, nil
}

func (r *SqlcCartsRepository) List() ([]entities.Cart, error) {
	rows, err := r.queries.ListCarts(r.ctx)
	if err != nil && err == pgx.ErrNoRows {
		return []entities.Cart{}, nil
	}

	if err != nil {
		return nil, err
	}

	carts := make([]entities.Cart, 0, len(rows))

	for _, row := range rows {
		cart := r.mapper.ToDomain(db.Cart(row))
		carts = append(carts, *cart)
	}

	return carts, nil
}

func (r *SqlcCartsRepository) Save(cart *entities.Cart) error {
	model := r.mapper.ToSqlc(*cart)

	params := db.UpdateCartStatusParams{}
	utils.Copy(&params, &model)

	return r.queries.UpdateCartStatus(r.ctx, params)
}

func (r *SqlcCartsRepository) DeleteById(id string) error {
	return r.queries.DeleteCartById(r.ctx, id)
}
