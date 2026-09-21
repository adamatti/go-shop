package repo

import "go.mongodb.org/mongo-driver/v2/bson"

// Mongo
type Product struct {
	ID                bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Name              string        `bson:"name" json:"name"`
	Description       string        `bson:"description" json:"description"`
	Price             float64       `bson:"price" json:"price"`
	AvailableQuantity int           `bson:"available_quantity" json:"availableQuantity"`
}

// Memory
type User struct {
	ID    string
	Email string
}

// Redis
type CardItem struct {
	ProductID string
	Quantity  int
}

type Card struct {
	UserID string
	Items  []CardItem
}
