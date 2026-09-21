package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

var startTime = time.Now()

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", healthHandler)
	mux.HandleFunc("GET /health", healthHandler)

	mux.HandleFunc("GET /api/products", listProductsHandler)
	mux.HandleFunc("POST /api/products", insertProductHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := ":" + port
	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
