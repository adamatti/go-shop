package main

import (
	"log"

	"adamatti.github.io/go-shop/repo"
	"github.com/spf13/cobra"
)

var cmdInsertProduct = &cobra.Command{
	Use:     "insert-product",
	Aliases: []string{"ip"},
	Short:   "Insert a sample product into the database",
	Run: func(cmd *cobra.Command, args []string) {
		err := repo.InsertProduct(repo.Product{
			Name:              "Sample Product",
			Description:       "This is a sample product.",
			Price:             19.99,
			AvailableQuantity: 100,
		})

		if err != nil {
			log.Fatalf("Failed to insert product: %v", err)
		}

		log.Println("Product inserted successfully")
	},
}

var cmdListProducts = &cobra.Command{
	Use:     "list-products",
	Aliases: []string{"lp"},
	Short:   "List all products in the database",
	Run: func(cmd *cobra.Command, args []string) {
		products, err := repo.ListProducts()
		if err != nil {
			log.Fatalf("Failed to list products: %v", err)
		}

		for _, product := range products {
			log.Printf("ID: %s, Name: %s, Description: %s, Price: %.2f, Available Quantity: %d",
				product.ID, product.Name, product.Description, product.Price, product.AvailableQuantity)
		}
	},
}
