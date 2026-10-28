package transactions

import (
	"github.com/gatimugabriel/1m-transactions/internal/db"
)

func updateDatabase(tx Transaction) error {
	query := "INSERT INTO transactions (order_id, user_id, product_id, quantity, price, timestamp) VALUES ($1, $2, $3, $4, $5, $6) RETURNING order_id"

	var orderId int

	err := db.DB.QueryRow(query, tx.OrderId, tx.UserId, tx.ProductId, tx.Quantity, tx.Price, tx.Timestamp).Scan(&orderId)
	if err != nil {
		return err
	}
	return nil
}
