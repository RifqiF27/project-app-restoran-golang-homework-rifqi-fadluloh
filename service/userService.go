package service

import (
	"database/sql"
	"fmt"
	// "os"
	"restaurant-app/model"
	"restaurant-app/repository"
	"restaurant-app/utils"
)

type UserService struct {
	RepoUser repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{RepoUser: repo}
}

func GetAllUsers(db *sql.DB) error {
	userRepo := repository.NewUserRepo(db)
	var users []model.User

	err := userRepo.GetAll(&users)
	if err != nil {
		return fmt.Errorf("gagal mengambil data user: %w", err)
	}

	for _, user := range users {
		fmt.Printf("ID: %d, Username: %s, Role: %s\n", user.ID, user.Username, user.Role)
	}

	return nil
}

func (us *UserService) LoginService(user model.User) (*model.User, error) {
	valid := utils.ReadLoggedIn("session.json")
	if valid {
		return nil,nil
	}
	
	err := utils.ReadBodyJSON("body.json", &user)
	if err != nil {
		return nil, nil
	}

	if user.Username == "" {
		utils.SendJSONResponse(400, "username cannot empty", nil)
		return nil, nil
	}
	if user.Password == "" {
		utils.SendJSONResponse(400, "password cannot empty", nil)
		return nil, nil
	}

	users, err := us.RepoUser.GetUserLogin(user)
	if err != nil {
		utils.SendJSONResponse(400, "username or password invalid", nil)
	} else {
		utils.SendJSONResponse(200, "login success", users)
		sessionData := map[string]interface{}{
			"ID":       users.ID,
			"Username": users.Username,
			"Role":     users.Role,
		}

		err = utils.WriteJSONFile("session.json", sessionData)
		if err != nil {
			utils.SendJSONResponse(500, err.Error(),nil)

		}

	}
	return users, nil

}
