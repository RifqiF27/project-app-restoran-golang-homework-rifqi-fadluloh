package service

import (
	"restaurant-app/model"
	"restaurant-app/repository"
	"restaurant-app/utils"
)

type RatingService struct {
	Repo      *repository.RatingRepository
	OrderRepo *repository.OrderRepository
}

func NewRatingService(repo *repository.RatingRepository, orderRepo *repository.OrderRepository) *RatingService {
	return &RatingService{Repo: repo, OrderRepo: orderRepo}
}

func (r *RatingService) AddRating(rating *model.Rating) error {
	userID, role, isAuthorized := utils.SessionRole()
	if !isAuthorized {
		return nil
	}
	if role != "Customer" {
		utils.SendJSONResponse(403, "Forbidden: only Customer can give rating", nil)

		return nil
	}
	err := utils.ReadBodyJSON("body.json", &rating)
	if err != nil {
		return nil
	}

	rating.CustomerID = userID

	existingOrder, err := r.OrderRepo.GetOrderByID(rating.OrderID)
	if err != nil {
		utils.SendJSONResponse(404, "Order id not found", nil)
		return err
	}

	if existingOrder.Status != "Completed" {
		utils.SendJSONResponse(400, "can't give rating if cancel order or order status is not completed", nil)
		return nil
	}
	ratingExists, err := r.Repo.RatingExistsForOrder(rating.CustomerID, rating.OrderID)
	if err != nil {
		utils.SendJSONResponse(500, err.Error(), nil)
		return nil
	}
	if ratingExists {
		utils.SendJSONResponse(400, "You have already rated this order", nil)
		return nil
	}

	if rating.Rating < 1 || rating.Rating > 5 {
		utils.SendJSONResponse(400, "rating must be between 1 and 5", nil)
		return nil
	}
	err = r.Repo.Create(rating)
	if err != nil {
		utils.SendJSONResponse(500, err.Error(), nil)

		return nil
	}
	utils.SendJSONResponse(201, "rating added successfully", rating)
	return nil

}

func (s *RatingService) GetRatings() ([]model.Rating, error) {
	userID, role, isAuthorized := utils.SessionRole()
	if !isAuthorized {
		return nil, nil
	}

	var ratings []model.Rating
	var err error
	if role == "Admin" || role == "Chef" {
		ratings, err = s.Repo.GetRatingsAll()

	} else if role == "Customer" {
		ratings, err = s.Repo.GetRatingsByOrderID(userID)

	}

	if err != nil {
		utils.SendJSONResponse(500, err.Error(), nil)
		return nil, err
	}
	utils.SendJSONResponse(200, "Ratings retrieved successfully", ratings)

	return ratings, nil
}
