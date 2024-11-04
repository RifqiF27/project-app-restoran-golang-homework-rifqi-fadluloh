package handler

import (
	"database/sql"
	"restaurant-app/model"
	"restaurant-app/repository"
	"restaurant-app/service"
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

	orderRepo := repository.NewOrderRepository(db)
	userRepo := repository.NewUserRepo(db)
	discountRepo := repository.NewDiscountRepository(db)
	menuRepo := repository.NewMenuItemRepository(db)

	orderService := service.NewOrderService(orderRepo, discountRepo, userRepo, menuRepo)
	orderService.UpdateStatus(&order)

}

func GetOrderItems(db *sql.DB) {

	orderRepo := repository.NewOrderRepository(db)
	userRepo := repository.NewUserRepo(db)
	discountRepo := repository.NewDiscountRepository(db)
	menuRepo := repository.NewMenuItemRepository(db)

	orderService := service.NewOrderService(orderRepo, discountRepo, userRepo, menuRepo)
	orderService.GetOrderHistory()

}

func DeleteOrder(db *sql.DB) {
	var order model.Order

	orderRepo := repository.NewOrderRepository(db)
	userRepo := repository.NewUserRepo(db)
	discountRepo := repository.NewDiscountRepository(db)
	menuRepo := repository.NewMenuItemRepository(db)

	orderService := service.NewOrderService(orderRepo, discountRepo, userRepo, menuRepo)
	orderService.DeleteOrder(&order)
	
}
