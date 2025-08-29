package repositories

import (
	"context"
	"errors"

	"github.com/natasha-m-oliveira/clean-architecture-go/adapter/database/prisma/mappers"
	"github.com/natasha-m-oliveira/clean-architecture-go/adapter/utils"
	"github.com/natasha-m-oliveira/clean-architecture-go/core/entities"
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

	_, err := r.client.Cart.CreateOne(status, optional...).Exec(r.ctx)
	if err != nil {
		return err
	}

	itemsTransactions := r.createItemsTransaction(cart.Items)

	err = r.client.Prisma.Transaction(itemsTransactions...).Exec(r.ctx)
	if err != nil {
		r.client.Cart.FindUnique(db.Cart.ID.Equals(cart.Id)).Delete().Exec(r.ctx)
	}

	return err
}

func (r *PrismaCartsRepository) createItemsTransaction(items []entities.CartItem) []transaction.Transaction {
	transactions := make([]transaction.Transaction, 0, len(items))

	for _, item := range items {
		quantity, cart, product, optional := (mappers.PrismaCartItemMapper{}).ToPrisma(item)
		tx := r.client.CartItem.CreateOne(quantity, cart, product, optional...).Tx()
		transactions = append(transactions, tx)
	}

	return transactions
}

func (r *PrismaCartsRepository) FindById(id string) (*entities.Cart, error) {
	model, err := r.client.Cart.FindUnique(db.Cart.ID.Equals(id)).Exec(r.ctx)
	if err != nil && errors.Is(err, db.ErrNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	cart := (mappers.PrismaCartMapper{}).ToDomain(*model)

	return &cart, err
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

	return err
}

func (r *PrismaCartsRepository) DeleteById(id string) error {
	_, err := r.client.Cart.FindUnique(db.Cart.ID.Equals(id)).Delete().Exec(r.ctx)

	return err
}
