package output

type Cart struct {
	Id     string `json:"id"`
	Status string `json:"status"`

	Items []CartItem `json:"items"`
}
