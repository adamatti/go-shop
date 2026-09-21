package main

import (
	"log"
	"net/http"
	"os"
)

func startHttp() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", healthHandler)
	mux.HandleFunc("GET /health", healthHandler)

	mux.HandleFunc("GET /api/products", listProductsHandler)
	mux.HandleFunc("POST /api/products", insertProductHandler)

	mux.HandleFunc("POST /api/cart", addToCartHandler)
	mux.HandleFunc("POST /api/cart/items", addToCartHandler)
	mux.HandleFunc("GET /api/cart", getCartHandler)
	mux.HandleFunc("DELETE /api/cart", deleteCartHandler)

	// TODO: use a config object
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := ":" + port
	log.Printf("listening on %s", addr)
	handler := fakeSessionMiddleware(mux)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatal(err)
	}
}
