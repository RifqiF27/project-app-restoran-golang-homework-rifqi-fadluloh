package handler

import (
	"database/sql"

	"restaurant-app/model"
	"restaurant-app/repository"
	"restaurant-app/service"
)

func Login(db *sql.DB) {
	user := model.User{}

	repo := repository.NewUserRepo(db)
	adminService := service.NewUserService(repo)

	adminService.LoginService(user)

}
