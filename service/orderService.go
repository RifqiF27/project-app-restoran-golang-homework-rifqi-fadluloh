package service

import (
	"restaurant-app/model"
	"restaurant-app/repository"
	"restaurant-app/utils"
)

type OrderService struct {
	OrderRepo    *repository.OrderRepository
	DiscountRepo *repository.DiscountRepository
	UserRepo     repository.UserRepository
	MenuItemRepo *repository.MenuItemRepository
}

func NewOrderService(orderRepo *repository.OrderRepository, discountRepo *repository.DiscountRepository, userRepo repository.UserRepository, menuRepo *repository.MenuItemRepository) *OrderService {
	return &OrderService{
		OrderRepo:    orderRepo,
		DiscountRepo: discountRepo,
		UserRepo:     userRepo,
		MenuItemRepo: menuRepo,
	}
}

func (s *OrderService) CreateOrder(order *model.Order) error {
    _, valid := utils.SessionAdmin()
	if !valid {
		return nil
	}
	err := utils.ReadBodyJSON("body.json", &order)
	if err != nil {
		return nil
	}

	customer, err := s.UserRepo.GetUserByUsername(order.CustomerID)
	if err != nil || customer == nil {
		utils.SendJSONResponse(400, "invalid customer ID", nil)
        return nil
	}

	var totalPrice float64

	for i := range order.Items {
		menuItem, err := s.MenuItemRepo.GetMenuItemByID(order.Items[i].MenuItemID)
		if err != nil {
			utils.SendJSONResponse(403, "menu not found", nil)
            return nil
		}

		totalPrice += menuItem.Price * float64(order.Items[i].Quantity)
		order.Items[i].Price = totalPrice
	}

	order.TotalPrice = totalPrice
	if order.DiscountID != 0 {
		discount, err := s.DiscountRepo.GetDiscountByCode(order.DiscountID)
		if err != nil {
			utils.SendJSONResponse(400, "invalid discount code", nil)
            return nil
		}
		order.TotalPrice = order.TotalPrice * (1 - discount.Percentage/100)
	}

	order.Status = "Pending"
	err = s.OrderRepo.CreateOrder(order)
	if err != nil {
		utils.SendJSONResponse(500, err.Error(), nil)

		return nil
	}
	utils.SendJSONResponse(200, "Order created successfully", order)

	return nil
}

func (s *OrderService) GetOrder(orderID int) (*model.Order, error) {
	return s.OrderRepo.GetOrderByID(orderID)
}
func (s *OrderService) UpdateStatus(id uint16, status string)  error {
	// var order model.Order
	// err := utils.ReadBodyJSON("body.json", &order)
	// if err != nil {
	// 	return nil
	// }
	validStatuses := []string{"Pending", "Processing", "Completed", "Canceled"}
    isValid := false
    for _, newStatus := range validStatuses {
        if status == newStatus {
            isValid = true
            break
        }
    }
    if !isValid {
		utils.SendJSONResponse(400, "invalid status", nil)
        return nil
    }
	err := s.OrderRepo.UpdateStatus(id, status)
	if err != nil {
		utils.SendJSONResponse(500, err.Error(), nil)
		return nil
	}
	utils.SendJSONResponse(200, "Status updated successfully",nil)
	return nil
}

