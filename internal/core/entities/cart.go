package entities

import "time"

const (
	CartStatusPending CartStatus = "pending"
	CartStatusOrdered CartStatus = "ordered"
)

type CartStatus string

type Cart struct {
	Id        string
	Status    CartStatus
	CreatedAt time.Time
	UpdatedAt time.Time

	Items []CartItem
}

type CartOptions struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCart(status CartStatus, items []CartItem, options ...CartOptions) *Cart {
	var id string
	var createdAt, updatedAt time.Time

	if len(options) > 0 {
		id = options[0].Id
		createdAt = options[0].CreatedAt
		updatedAt = options[0].UpdatedAt
	}

	if createdAt.IsZero() {
		createdAt = time.Now()
	}

	if updatedAt.IsZero() {
		updatedAt = createdAt
	}

	return &Cart{
		Id:        id,
		Status:    status,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Items:     items,
	}
}
