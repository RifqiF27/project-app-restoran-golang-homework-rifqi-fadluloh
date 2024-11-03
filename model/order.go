package model

import "time"

type Order struct {
	ID         int         `json:"id"`
	CustomerID int         `json:"customer_id"`
	Status     string      `json:"status"`
	DiscountID int         `json:"discount_id"`
	TotalPrice float64     `json:"total_price"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	Items      []OrderItem `json:"items"`
}
