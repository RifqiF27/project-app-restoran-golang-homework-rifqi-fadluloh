package repository

import (
	"database/sql"
	"restaurant-app/model"
	"time"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (repo *OrderRepository) CreateOrder(order *model.Order) error {
	query := `INSERT INTO "Orders" (customer_id, status, discount_id, total_price, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := repo.DB.QueryRow(query, order.CustomerID, order.Status, order.DiscountID, order.TotalPrice, time.Now(), time.Now()).Scan(&order.ID)
	if err != nil {
		return err
	}

	for i, item := range order.Items {
		itemQuery := `INSERT INTO "Order_Items" (order_id, menu_item_id, quantity, price) 
                      VALUES ($1, $2, $3, $4) RETURNING id`
		err := repo.DB.QueryRow(itemQuery, order.ID, item.MenuItemID, item.Quantity, item.Price).Scan(&order.Items[i].ID)
		if err != nil {
			return err
		}

		order.Items[i].OrderID = order.ID
	}
	return nil
}

func (repo *OrderRepository) GetOrderByID(orderID int) (*model.Order, error) {
	order := &model.Order{}
	query := `SELECT id, customer_id, status, discount_id, total_price FROM "Orders" WHERE id = $1`
	err := repo.DB.QueryRow(query, orderID).Scan(&order.ID, &order.CustomerID, &order.Status, &order.DiscountID, &order.TotalPrice)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (repo *OrderRepository) UpdateStatus(id uint16, status string) error {
	query := `UPDATE "Orders" SET status = $1 WHERE id = $2`
	_, err := repo.DB.Exec(query, status, id)
	return err
}
