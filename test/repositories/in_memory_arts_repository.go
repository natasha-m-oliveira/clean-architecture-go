package repositories

import (
	"sync"

	"github.com/google/uuid"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/errors"
)

type InMemoryCartsRepository struct {
	mu    sync.Mutex
	carts map[string]*entities.Cart
}

func NewInMemoryCartsRepository() *InMemoryCartsRepository {
	return &InMemoryCartsRepository{
		carts: make(map[string]*entities.Cart),
	}
}

func (r *InMemoryCartsRepository) Create(cart *entities.Cart) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	cart.Id = uuid.New().String()

	r.carts[cart.Id] = cart
	return nil
}

func (r *InMemoryCartsRepository) FindById(id string) (*entities.Cart, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	cart, ok := r.carts[id]
	if !ok {
		return nil, &errors.CartNotFound{}
	}
	return cart, nil
}

func (r *InMemoryCartsRepository) Save(cart *entities.Cart) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.carts[cart.Id] = cart
	return nil
}

func (r *InMemoryCartsRepository) DeleteById(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.carts, id)
	return nil
}
