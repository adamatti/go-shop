package main

import (
	"context"
	"log"
	"time"

	pb "adamatti.github.io/go-shop/grpc"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const grpcAddress = "localhost:50051"

func getGrpcClient() (*grpc.ClientConn, pb.ShopClient) {
	conn, err := grpc.NewClient(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	return conn, pb.NewShopClient(conn)
}

var cmdGrpcInsertProduct = &cobra.Command{
	Use:     "grpc-insert-product",
	Aliases: []string{"gip"},
	Short:   "Insert a sample product via gRPC",
	Run: func(cmd *cobra.Command, args []string) {
		conn, client := getGrpcClient()
		defer func() { _ = conn.Close() }()

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		resp, err := client.InsertProduct(ctx, &pb.InsertProductRequest{
			Product: &pb.Product{
				Name:              "Sample Product",
				Description:       "This is a sample product.",
				Price:             19.99,
				AvailableQuantity: 100,
			},
		})
		if err != nil {
			log.Fatalf("Failed to insert product: %v", err)
		}

		log.Printf("Product inserted successfully: %s", resp.GetId())
	},
}

var cmdGrpcListProducts = &cobra.Command{
	Use:     "grpc-list-products",
	Aliases: []string{"glp"},
	Short:   "List all products via gRPC",
	Run: func(cmd *cobra.Command, args []string) {
		conn, client := getGrpcClient()
		defer func() { _ = conn.Close() }()

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		resp, err := client.ListProducts(ctx, &pb.ListProductsRequest{})
		if err != nil {
			log.Fatalf("Failed to list products: %v", err)
		}

		for _, product := range resp.GetProducts() {
			log.Printf("ID: %s, Name: %s, Description: %s, Price: %.2f, Available Quantity: %d",
				product.GetId(), product.GetName(), product.GetDescription(), product.GetPrice(), product.GetAvailableQuantity())
		}
	},
}
