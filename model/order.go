package model

import "time"

type Order struct {
	ID         int         `json:"id,omitempty"`
	CustomerID int         `json:"customer_id,omitempty"`
	Status     string      `json:"status,omitempty"`
	DiscountID int         `json:"discount_id,omitempty"`
	TotalPrice float64     `json:"total_price,omitempty"`
	CreatedAt  time.Time   `json:"created_at,omitempty"`
	UpdatedAt  time.Time   `json:"updated_at,omitempty"`
	Items      []OrderItem `json:"items,omitempty"`
}
