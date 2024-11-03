package repository

import (
	"database/sql"
	"restaurant-app/model"
)

type DiscountRepository struct {
	DB *sql.DB
}

func NewDiscountRepository(db *sql.DB) *DiscountRepository {
	return &DiscountRepository{DB: db}
}

func (repo *DiscountRepository) GetDiscountByCode(id int) (*model.Discount, error) {
	discount := &model.Discount{}
	query := `SELECT id, code, percentage FROM "Discounts" WHERE id = $1`
	err := repo.DB.QueryRow(query, id).Scan(&discount.ID, &discount.Code, &discount.Percentage)
	if err != nil {
		return nil, err
	}
	return discount, nil
}
