package main

import (
	"encoding/json"
	"net/http"

	"adamatti.github.io/go-shop/service"
)

type OrderRequest struct {
	OrderID string `json:"orderId" validate:"required"`
}

func submitOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	req := &OrderRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(req); err != nil {
		processError(err, w)
		return
	}

	userID := getUserIDFromSession(r)
	if userID == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	if err := service.SubmitOrder(userID, req.OrderID); err != nil {
		processError(err, w)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
