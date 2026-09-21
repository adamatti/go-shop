package service

import (
	"encoding/json"
	"fmt"
	"strings"

	"adamatti.github.io/go-shop/repo"
)

func cartKey(userID string) string {
	return fmt.Sprintf("cart:%s", userID)
}

// AddToCart adds an item to the user's cart.
// If the product already exists in the cart, its quantity is incremented.
func AddToCart(userID string, item CartItem) error {
	userID = strings.TrimSpace(userID)
	if userID == "" {
		return ErrInvalidUserID
	}
	if strings.TrimSpace(item.ProductID) == "" {
		return ErrInvalidProductID
	}
	if item.Quantity <= 0 {
		return ErrInvalidQuantity
	}

	cart, err := GetCart(userID)
	if err != nil {
		return err
	}

	updated := false
	for i, it := range cart.Items {
		if it.ProductID == item.ProductID {
			cart.Items[i].Quantity += item.Quantity
			updated = true
			break
		}
	}

	if !updated {
		cart.Items = append(cart.Items, item)
	}

	return saveCart(cart)
}

// GetCart retrieves the cart for a user.
// Returns an empty cart if none exists yet.
func GetCart(userID string) (*Cart, error) {
	userID = strings.TrimSpace(userID)
	if userID == "" {
		return nil, ErrInvalidUserID
	}

	data, err := repo.Get(cartKey(userID))
	if err != nil {
		return nil, err
	}

	if data == nil {
		return &Cart{UserID: userID, Items: []CartItem{}}, nil
	}

	var cart Cart
	if err := json.Unmarshal(data, &cart); err != nil {
		return nil, fmt.Errorf("failed to unmarshal cart: %w", err)
	}
	cart.UserID = userID

	if cart.Items == nil {
		cart.Items = []CartItem{}
	}

	return &cart, nil
}

// DeleteCart removes the user's cart entirely.
func DeleteCart(userID string) error {
	userID = strings.TrimSpace(userID)
	if userID == "" {
		return ErrInvalidUserID
	}

	return repo.Delete(cartKey(userID))
}

func saveCart(cart *Cart) error {
	data, err := json.Marshal(cart)
	if err != nil {
		return fmt.Errorf("failed to marshal cart: %w", err)
	}

	return repo.Set(cartKey(cart.UserID), data)
}
