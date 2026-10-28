package transactions

import "time"

type Transaction struct {
	OrderId   int       `json:"order_id"`
	UserId    string    `json:"user_id"`
	ProductId string    `json:"product_id"`
	Quantity  int32     `json:"quantity"`
	Price     int64     `json:"price"` // in cents
	Timestamp time.Time `json:"timestamp"`
}
