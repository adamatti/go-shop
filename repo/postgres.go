package repo

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresClient *gorm.DB

func init() {
	dsn := "host=localhost user=postgres password=postgres dbname=app port=5432 sslmode=disable"
	var err error
	PostgresClient, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// For demostration only, to avoid migrations
	if err := PostgresClient.AutoMigrate(&Order{}); err != nil {
		log.Fatal(err)
	}
	if err := PostgresClient.AutoMigrate(&OrderItem{}); err != nil {
		log.Fatal(err)
	}
}
