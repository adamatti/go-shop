package service

// CartItem represents a single product entry in a user's cart.
type CartItem struct {
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

// Cart represents a user's shopping cart.
type Cart struct {
	UserID string     `json:"-"`
	Items  []CartItem `json:"items"`
}
