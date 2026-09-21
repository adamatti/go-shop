package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"adamatti.github.io/go-shop/service"
)

type addToCartRequest struct {
	ProductID string `json:"productId" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required,gt=0"`
}

func addToCartHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := getUserIDFromSession(r)
	if userID == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var req addToCartRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		processError(err, w)
		return
	}

	if err := validate.Struct(req); err != nil {
		processError(err, w)
		return
	}

	item := service.CartItem{
		ProductID: strings.TrimSpace(req.ProductID),
		Quantity:  req.Quantity,
	}

	if err := service.AddToCart(userID, item); err != nil {
		http.Error(w, "failed to add item to cart", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func getCartHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := getUserIDFromSession(r)
	if userID == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	cart, err := service.GetCart(userID)
	if err != nil {
		http.Error(w, "failed to get cart", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(cart)
}

func deleteCartHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := getUserIDFromSession(r)
	if userID == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	if err := service.DeleteCart(userID); err != nil {
		http.Error(w, "failed to delete cart", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
