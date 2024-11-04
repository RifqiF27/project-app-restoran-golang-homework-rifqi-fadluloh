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

	userID, role, isAuthorized := utils.SessionRole()
	if !isAuthorized {
		return nil
	}
	if role != "Customer" {
		utils.SendJSONResponse(403, "Forbidden: only Customer can order", nil)

		return nil
	}
	err := utils.ReadBodyJSON("body.json", &order)
	if err != nil {
		return nil
	}

	order.CustomerID = userID

	var totalPrice float64
	if len(order.Items) <= 0 {
		utils.SendJSONResponse(400, "items cannot empty", nil)
		return nil
	}
	for i := range order.Items {
		menuItem, err := s.MenuItemRepo.GetMenuItemByID(order.Items[i].MenuItemID)
		if err != nil {
			utils.SendJSONResponse(404, "menu not found", nil)
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

func (s *OrderService) UpdateStatus(order *model.Order) error {
	_, role, isAuthorized := utils.SessionRole()
	if !isAuthorized {
		return nil
	}
	if role != "Admin" && role != "Chef" {
		utils.SendJSONResponse(403, "Forbidden: only Admin and Chef can update order status", nil)

		return nil
	}
	err := utils.ReadBodyJSON("body.json", &order)
	if err != nil {
		return nil
	}
	existingOrder, err := s.OrderRepo.GetOrderByID(order.ID)
	if err != nil {
		utils.SendJSONResponse(404, "Order id not found", nil)
		return err
	}

	if existingOrder.Status == "Completed" || existingOrder.Status == "Canceled" {
		utils.SendJSONResponse(400, "status cannot be updated as it is already completed or canceled", nil)
		return nil
	}
	validStatuses := []string{"Pending", "Processing", "Completed", "Canceled"}
	isValid := false
	for _, newStatus := range validStatuses {
		if order.Status == newStatus {
			isValid = true
			break
		}
	}
	if !isValid {
		utils.SendJSONResponse(400, "invalid status", nil)
		return nil
	}
	err = s.OrderRepo.UpdateStatus(uint16(order.ID), order.Status)
	if err != nil {
		utils.SendJSONResponse(500, err.Error(), nil)
		return nil
	}
	utils.SendJSONResponse(200, "Status updated successfully", order)
	return nil
}

func (s *OrderService) GetOrderHistory() ([]model.Order, error) {
	userID, role, isAuthorized := utils.SessionRole()
	if !isAuthorized {
		return nil, nil
	}

	var orders []model.Order
	var err error

	switch role {
	case "Admin":
		orders, err = s.OrderRepo.GetAllOrdersWithItems()
	case "Chef":
		orders, err = s.OrderRepo.GetAllOrdersWithItems()
	case "Customer":
		orders, err = s.OrderRepo.GetOrdersByCustomerID(userID)
	default:
		utils.SendJSONResponse(403, "Forbidden", nil)
		return nil, nil
	}

	if err != nil {
		utils.SendJSONResponse(500, err.Error(), nil)
		return nil, err
	}

	utils.SendJSONResponse(200, "Order history retrieved successfully", orders)
	return orders, nil
}

func (s *OrderService) DeleteOrder(order *model.Order) error {

	_, role, isAuthorized := utils.SessionRole()
	if !isAuthorized {
		return nil
	}
	if role != "Admin" {
		utils.SendJSONResponse(403, "Forbidden: only Admin can delete orders", nil)
		return nil
	}
	err := utils.ReadBodyJSON("body.json", &order)
	if err != nil {
		return nil
	}
	if order.ID == 0 {
		utils.SendJSONResponse(400, "Order ID cannot empty", nil)
		return nil
	}
	existingOrder, err := s.OrderRepo.GetOrderByID(order.ID)
	if err != nil || existingOrder.ID == 0 {
		utils.SendJSONResponse(404, "Order id not found", nil)
		return err
	}

	err = s.OrderRepo.DeleteOrderByID(order.ID)
	if err != nil {
		utils.SendJSONResponse(500, err.Error(), nil)
		return err
	}

	utils.SendJSONResponse(200, "Order deleted successfully", nil)
	return nil
}
