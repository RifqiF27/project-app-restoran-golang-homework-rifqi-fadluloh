package handler

import (
	"database/sql"
	"restaurant-app/model"
	"restaurant-app/repository"
	"restaurant-app/service"
)

func AddMenu(db *sql.DB) {
	var menu model.MenuItem

	menuRepo := repository.NewMenuItemRepository(db)

	menuService := service.NewMenuItemService(*menuRepo)
	menuService.AddMenuService(&menu)
}
