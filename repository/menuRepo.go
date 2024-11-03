package repository

import (
	"database/sql"
	"restaurant-app/model"
)

type MenuItemRepository struct {
	DB *sql.DB
}

func NewMenuItemRepository(db *sql.DB) *MenuItemRepository {
	return &MenuItemRepository{DB: db}
}

func (repo *MenuItemRepository) GetMenuItemByID(menuItemID int) (*model.MenuItem, error) {
	menuItem := &model.MenuItem{}
	query := `SELECT id, name, price FROM "Menu_Items" WHERE id = $1`
	err := repo.DB.QueryRow(query, menuItemID).Scan(&menuItem.ID, &menuItem.Name, &menuItem.Price)
	if err != nil {
		return nil, err
	}
	return menuItem, nil
}
