package main

import (
	"encoding/json"
	"net/http"

	"adamatti.github.io/go-shop/repo"
)

func listProductsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// TODO: it should call service, not repo directly
	products, err := repo.ListProducts()
	if err != nil {
		http.Error(w, "failed to list products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(products)
}

func insertProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var product repo.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		processError(err, w)
		return
	}

	// TODO: it should call service, not repo directly
	if err := repo.InsertProduct(product); err != nil {
		http.Error(w, "failed to insert product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
