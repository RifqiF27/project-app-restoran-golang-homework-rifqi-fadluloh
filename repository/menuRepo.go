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

func (repo *MenuItemRepository) Create(menu *model.MenuItem) error {
	query := `INSERT INTO "Menu_Items" (name, price) VALUES ($1, $2) RETURNING id`
	err := repo.DB.QueryRow(query, menu.Name, menu.Price).Scan(&menu.ID)
	if err != nil {
		return err
	}

	return nil
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
