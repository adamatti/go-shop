package service

import "errors"

var (
	ErrInvalidUserID    = errors.New("user id cannot be empty")
	ErrInvalidOrderID   = errors.New("order id cannot be empty")
	ErrInvalidProductID = errors.New("product id cannot be empty")
	ErrInvalidQuantity  = errors.New("quantity must be greater than 0")
	ErrEmptyCart        = errors.New("cart is empty")
)
