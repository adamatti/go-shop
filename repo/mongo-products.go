package repo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const (
	// TODO: Move the connection string to a config file or environment variable
	databaseURI  = "mongodb://localhost:27017"
	databaseName = "go-shop"
)

var client *mongo.Client
var productCollection *mongo.Collection

func init() {
	var err error
	client, err = mongo.Connect(options.Client().ApplyURI(databaseURI))
	if err != nil {
		log.Fatal(err)
	}
	productCollection = client.Database(databaseName).Collection("products")
}

func InsertProduct(product Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := productCollection.InsertOne(ctx, product)
	if err != nil {
		return err
	}
	log.Println("product inserted:", res.InsertedID)
	return nil
}

// Note: service layer should translate DB model to business model
func ListProducts() ([]Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := productCollection.Find(ctx, map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx) //nolint:errcheck

	var products []Product
	for cursor.Next(ctx) {
		var product Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
