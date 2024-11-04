package repository

import (
	"database/sql"
	"restaurant-app/model"
)

type UserRepository interface {
	GetUserLogin(user model.User) (*model.User, error)
}

type UserRepoDb struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepository {
	return &UserRepoDb{DB: db}
}

func (r *UserRepoDb) GetUserLogin(user model.User) (*model.User, error) {
	query := `SELECT id, username, password, role FROM "Users" WHERE username=$1 AND password=$2`
	var userResponse model.User
	err := r.DB.QueryRow(query, user.Username, user.Password).Scan(&userResponse.ID, &userResponse.Username, &userResponse.Password, &userResponse.Role)

	if err != nil {
		return nil, err
	}

	return &userResponse, nil
}
