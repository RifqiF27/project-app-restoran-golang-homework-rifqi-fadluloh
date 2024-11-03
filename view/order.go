package handler

import (
	"database/sql"
	"restaurant-app/model"
	"restaurant-app/repository"
	"restaurant-app/service"
	"restaurant-app/utils"
)

func AddOrder(db *sql.DB) {

	var order model.Order

	orderRepo := repository.NewOrderRepository(db)
	userRepo := repository.NewUserRepo(db)
	discountRepo := repository.NewDiscountRepository(db)
	menuRepo := repository.NewMenuItemRepository(db)

	orderService := service.NewOrderService(orderRepo, discountRepo, userRepo, menuRepo)

	if err := orderService.CreateOrder(&order); err != nil {
		return
	}
}

func UpdateStatus(db *sql.DB) {
	var order model.Order
	err := utils.ReadBodyJSON("body.json", &order)
	if err != nil {
		return 
	}

	orderRepo := repository.NewOrderRepository(db)
	userRepo := repository.NewUserRepo(db)
	discountRepo := repository.NewDiscountRepository(db)
	menuRepo := repository.NewMenuItemRepository(db)

	orderService := service.NewOrderService(orderRepo, discountRepo, userRepo, menuRepo)
	err = orderService.UpdateStatus(uint16(order.ID), order.Status)
	if err != nil {
		return 
	} 
	
	
}
