package output

import (
	"time"
)

type Cart struct {
	Id        string    `json:"id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Items []CartItem `json:"items"`
}
