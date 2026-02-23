package repositories

import (
	"maps"
	"slices"
	"sync"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/entities"
)

type InMemoryProductsRepository struct {
	mu       sync.Mutex
	products map[string]entities.Product
}

func NewInMemoryProductsRepository() *InMemoryProductsRepository {
	return &InMemoryProductsRepository{
		products: make(map[string]entities.Product),
	}
}

func (r *InMemoryProductsRepository) Create(product *entities.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.products[product.Id] = *product
	return nil
}

func (r *InMemoryProductsRepository) FindById(id string) (*entities.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	product, ok := r.products[id]
	if !ok {
		return nil, nil
	}
	return &product, nil
}

func (r *InMemoryProductsRepository) FindByName(name string) (*entities.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, product := range r.products {
		if product.Name == name {
			return &product, nil
		}
	}

	return nil, nil
}

func (r *InMemoryProductsRepository) FindAll() ([]entities.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	products := slices.Collect(maps.Values(r.products))

	return products, nil
}

func (r *InMemoryProductsRepository) Save(product *entities.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.products[product.Id] = *product
	return nil
}

func (r *InMemoryProductsRepository) DeleteById(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.products, id)
	return nil
}
