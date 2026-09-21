package repo

import (
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const (
	// TODO: Move the connection string to a config file or environment variable
	databaseURI  = "mongodb://localhost:27017"
	databaseName = "go-shop"
)

var mongoClient *mongo.Client
var productCollection *mongo.Collection

func init() {
	var err error
	mongoClient, err = mongo.Connect(options.Client().ApplyURI(databaseURI))
	if err != nil {
		log.Fatal(err)
	}
	productCollection = mongoClient.Database(databaseName).Collection("products")
}
