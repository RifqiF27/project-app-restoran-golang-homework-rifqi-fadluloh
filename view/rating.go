package handler

import (
	"database/sql"
	"restaurant-app/model"
	"restaurant-app/repository"
	"restaurant-app/service"
)

func AddRating(db *sql.DB) {
	var rating model.Rating

	ratingRepo := repository.NewRatingRepository(db)
	orderRepo := repository.NewOrderRepository(db)

	ratingService := service.NewRatingService(ratingRepo, orderRepo)

	ratingService.AddRating(&rating)

}

func GetRatings(db *sql.DB) {
	ratingRepo := repository.NewRatingRepository(db)
	orderRepo := repository.NewOrderRepository(db)

	ratingService := service.NewRatingService(ratingRepo,orderRepo)

	ratingService.GetRatings()
	
}
