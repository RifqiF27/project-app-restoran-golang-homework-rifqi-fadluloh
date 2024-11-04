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

func (repo *OrderRepository) GetAllOrdersWithItems() ([]model.Order, error) {

	query := `SELECT id, customer_id, status, discount_id, total_price, created_at, updated_at FROM "Orders"`
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []model.Order
	for rows.Next() {
		var order model.Order
		if err := rows.Scan(&order.ID, &order.CustomerID, &order.Status, &order.DiscountID, &order.TotalPrice, &order.CreatedAt, &order.UpdatedAt); err != nil {
			return nil, err
		}

		items, err := repo.GetOrderItemsByOrderID(order.ID)
		if err != nil {
			return nil, err
		}
		order.Items = items

		orders = append(orders, order)
	}

	return orders, nil
}

func (repo *OrderRepository) GetOrderItemsByOrderID(orderID int) ([]model.OrderItem, error) {
	query := `SELECT id, order_id, menu_item_id, quantity, price FROM "Order_Items" WHERE order_id = $1`
	rows, err := repo.DB.Query(query, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.OrderItem
	for rows.Next() {
		var item model.OrderItem
		if err := rows.Scan(&item.ID, &item.OrderID, &item.MenuItemID, &item.Quantity, &item.Price); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (repo *OrderRepository) GetOrdersByCustomerID(customerID int) ([]model.Order, error) {
	query := `SELECT id, customer_id, status, discount_id, total_price, created_at, updated_at FROM "Orders" WHERE customer_id = $1`
	rows, err := repo.DB.Query(query, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []model.Order
	for rows.Next() {
		var order model.Order
		if err := rows.Scan(&order.ID, &order.CustomerID, &order.Status, &order.DiscountID, &order.TotalPrice, &order.CreatedAt, &order.UpdatedAt); err != nil {
			return nil, err
		}

		items, err := repo.GetOrderItemsByOrderID(order.ID)
		if err != nil {
			return nil, err
		}
		order.Items = items

		orders = append(orders, order)
	}

	return orders, nil
}

func (repo *OrderRepository) DeleteOrderByID(orderID int) error {
	query := `DELETE FROM "Orders" WHERE id = $1`
	_, err := repo.DB.Exec(query, orderID)
	if err != nil {
		return err
	}
	return nil
}
