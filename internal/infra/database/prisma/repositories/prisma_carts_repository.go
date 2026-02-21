package repositories

import (
	"context"
	"errors"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/entities"
	_errors "github.com/natasha-m-oliveira/clean-architecture-go/internal/app/errors"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/database/prisma/repositories/mappers"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/utils"
	"github.com/natasha-m-oliveira/clean-architecture-go/prisma/db"
	"github.com/steebchen/prisma-client-go/runtime/transaction"
)

type PrismaCartsRepository struct {
	ctx        context.Context
	client     *db.PrismaClient
	cartMapper mappers.PrismaCartMapper
	itemMapper mappers.PrismaCartItemMapper
}

func NewPrismaCartsRepository(ctx context.Context, client *db.PrismaClient) *PrismaCartsRepository {
	return &PrismaCartsRepository{
		ctx:        ctx,
		client:     client,
		cartMapper: mappers.NewPrismaCartMapper(),
		itemMapper: mappers.NewPrismaCartItemMapper(),
	}
}

func (r *PrismaCartsRepository) Create(cart *entities.Cart) error {
	status, optional := r.cartMapper.ToPrisma(*cart)

	cartTransaction := r.client.Cart.CreateOne(status, optional...).Tx()
	itemsTransaction := r.createItemsTransaction(cart.Items)

	transaction := append(itemsTransaction, cartTransaction)

	err := r.client.Prisma.Transaction(transaction...).Exec(r.ctx)
	if err != nil {
		return _errors.NewInternalServerError()
	}

	return nil
}

func (r *PrismaCartsRepository) createItemsTransaction(items []entities.CartItem) []transaction.Transaction {
	transactions := make([]transaction.Transaction, len(items))

	for _, item := range items {
		quantity, cart, product, optional := r.itemMapper.ToPrisma(item)

		r.client.CartItem.CreateOne(quantity, cart, product, optional...)
	}

	return transactions
}

func (r *PrismaCartsRepository) FindById(id string) (*entities.Cart, error) {
	model, err := r.client.Cart.FindUnique(db.Cart.ID.Equals(id)).Exec(r.ctx)
	if err != nil && errors.Is(db.ErrNotFound, err) {
		return nil, nil
	}
	if err != nil {
		return nil, _errors.NewInternalServerError()
	}

	Cart := r.cartMapper.ToDomain(*model)

	return &Cart, err
}

func (r *PrismaCartsRepository) List() ([]entities.Cart, error) {
	models, err := r.client.Cart.FindMany().Exec(r.ctx)
	if err != nil {
		return []entities.Cart{}, err
	}

	Carts := utils.Map(models, r.cartMapper.ToDomain)

	return Carts, nil
}

func (r *PrismaCartsRepository) Save(Cart *entities.Cart) error {
	status, optional := r.cartMapper.ToPrisma(*Cart)

	optional = append(optional, status.(db.CartSetParam))

	_, err := r.client.Cart.FindUnique(db.Cart.ID.Equals(Cart.Id)).Update(optional...).Exec(r.ctx)
	if err != nil {
		return _errors.NewInternalServerError()
	}
	return nil
}

func (r *PrismaCartsRepository) DeleteById(id string) error {
	_, err := r.client.Cart.FindUnique(db.Cart.ID.Equals(id)).Delete().Exec(r.ctx)
	if err != nil {
		return _errors.NewInternalServerError()
	}
	return nil
}
