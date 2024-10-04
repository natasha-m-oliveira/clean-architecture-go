package repositories

import (
	"context"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/adapter/repositories/mappers"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/adapter/utils"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/errors"
	"github.com/natasha-m-oliveira/clean-architecture-go/prisma/db"
	"github.com/steebchen/prisma-client-go/runtime/transaction"
)

type PrismaCartsRepository struct {
	client *db.PrismaClient
	ctx    context.Context
}

func NewPrismaCartsRepository(client *db.PrismaClient, ctx context.Context) *PrismaCartsRepository {
	return &PrismaCartsRepository{
		client: client,
		ctx:    ctx,
	}
}

func (r *PrismaCartsRepository) Create(cart *entities.Cart) error {
	status, optional := (mappers.PrismaCartMapper{}).ToPrisma(*cart)

	cartTransaction := r.client.Cart.CreateOne(status, optional...).Tx()
	itemsTransaction := r.createItemsTransaction(cart.Items)

	transaction := append(itemsTransaction, cartTransaction)

	err := r.client.Prisma.Transaction(transaction...).Exec(r.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *PrismaCartsRepository) createItemsTransaction(items []entities.CartItem) []transaction.Transaction {
	transactions := make([]transaction.Transaction, len(items))

	for _, item := range items {
		quantity, cart, product, optional := (mappers.PrismaCartItemMapper{}).ToPrisma(item)

		r.client.CartItem.CreateOne(quantity, cart, product, optional...)
	}

	return transactions
}

func (r *PrismaCartsRepository) FindById(id string) (*entities.Cart, error) {
	model, err := r.client.Cart.FindUnique(db.Cart.ID.Equals(id)).Exec(r.ctx)
	if err != nil {
		return nil, utils.ParseError(err, errors.CartNotFound{})
	}

	Cart := (mappers.PrismaCartMapper{}).ToDomain(*model)

	return &Cart, err
}

func (r *PrismaCartsRepository) List() ([]entities.Cart, error) {
	models, err := r.client.Cart.FindMany().Exec(r.ctx)
	if err != nil {
		return []entities.Cart{}, err
	}

	Carts := utils.MapToDomain(models, (mappers.PrismaCartMapper{}).ToDomain)

	return Carts, nil
}

func (r *PrismaCartsRepository) Save(Cart *entities.Cart) error {
	status, optional := (mappers.PrismaCartMapper{}).ToPrisma(*Cart)

	optional = append(optional, status.(db.CartSetParam))

	_, err := r.client.Cart.FindUnique(db.Cart.ID.Equals(Cart.Id)).Update(optional...).Exec(r.ctx)

	return utils.ParseError(err, errors.CartNotFound{})
}

func (r *PrismaCartsRepository) DeleteById(id string) error {
	_, err := r.client.Cart.FindUnique(db.Cart.ID.Equals(id)).Delete().Exec(r.ctx)

	return utils.ParseError(err, errors.CartNotFound{})
}
