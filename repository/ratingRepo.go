package repository

import (
	"database/sql"
	"restaurant-app/model"
)

type RatingRepository struct {
	DB *sql.DB
}

func NewRatingRepository(db *sql.DB) *RatingRepository {
	return &RatingRepository{DB: db}
}

func (repo *RatingRepository) Create(rating *model.Rating) error {
	query := `INSERT INTO "Ratings" (order_id, customer_id, rating, comment) VALUES ($1, $2, $3, $4) RETURNING id`
	err := repo.DB.QueryRow(query, rating.OrderID, rating.CustomerID, rating.Rating, rating.Comment).Scan(&rating.ID)
	return err
}

func (repo *RatingRepository) GetRatingsAll() ([]model.Rating, error) {
	query := `SELECT id, order_id, customer_id, rating, comment FROM "Ratings"`
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ratings []model.Rating
	for rows.Next() {
		var rating model.Rating
		if err := rows.Scan(&rating.ID, &rating.OrderID, &rating.CustomerID, &rating.Rating, &rating.Comment); err != nil {
			return nil, err
		}
		ratings = append(ratings, rating)
	}
	return ratings, nil
}
func (repo *RatingRepository) GetRatingsByOrderID(CustomerID int) ([]model.Rating, error) {
	query := `SELECT id, order_id, customer_id, rating, comment FROM "Ratings" WHERE customer_id = $1`
	rows, err := repo.DB.Query(query, CustomerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ratings []model.Rating
	for rows.Next() {
		var rating model.Rating
		if err := rows.Scan(&rating.ID, &rating.OrderID, &rating.CustomerID, &rating.Rating, &rating.Comment); err != nil {
			return nil, err
		}
		ratings = append(ratings, rating)
	}
	return ratings, nil
}

func (repo *RatingRepository) RatingExistsForOrder(customerID, orderID int) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM "Ratings" WHERE customer_id = $1 AND order_id = $2)`
	var exists bool
	err := repo.DB.QueryRow(query, customerID, orderID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
