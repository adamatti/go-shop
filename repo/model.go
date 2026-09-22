package repo

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

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

// Postgres
type Order struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	UserId    string    `gorm:"not null;size:120" json:"userId"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

type OrderItem struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	OrderId   string    `gorm:"not null;size:120" json:"orderId"`
	Order     Order     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ProductId string    `gorm:"not null;size:120" json:"productId"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}
