package output

type CartItem struct {
	Id        string  `json:"id"`
	ProductId string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Product   Product `json:"product"`
}
