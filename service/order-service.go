package service

import (
	"strings"

	"adamatti.github.io/go-shop/repo"
	"gorm.io/gorm"
)

/*
 * Receiving order id here for idempotency
 */
func SubmitOrder(userID string, orderID string) error {
	userID = strings.TrimSpace(userID)
	if userID == "" {
		return ErrInvalidUserID
	}

	if orderID == "" {
		return ErrInvalidOrderID
	}
	cart, err := GetCart(userID)
	if err != nil {
		return err
	}
	if len(cart.Items) == 0 {
		return ErrEmptyCart
	}

	err = repo.PostgresClient.Transaction(func(tx *gorm.DB) error {
		// create order
		if err := tx.Create(&repo.Order{UserId: userID, ID: orderID}).Error; err != nil {
			return err
		}
		// create order items
		for _, item := range cart.Items {
			if err := tx.Create(&repo.OrderItem{OrderId: orderID, ProductId: item.ProductID, Quantity: item.Quantity}).Error; err != nil {
				return err
			}
		}

		// submit order to queue
		if err := SendToKafka("orders", []byte(orderID), map[string]interface{}{
			"order_id": orderID,
			"user_id":  userID,
		}); err != nil {
			return err
		}

		return nil // commit
	})

	if err != nil {
		return err
	}

	err = DeleteCart(userID)
	if err != nil {
		return err
	}

	// TODO submit to queue here

	return nil
}
