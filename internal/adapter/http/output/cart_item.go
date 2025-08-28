package output

import "time"

type CartItem struct {
	Id        string    `json:"id"`
	CartId    string    `json:"cart_id"`
	ProductId string    `json:"product_id"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`

	Product Product `json:"product,omitempty"`
}
