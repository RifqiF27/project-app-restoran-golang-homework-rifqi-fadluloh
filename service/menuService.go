package service

import (
	"restaurant-app/model"
	"restaurant-app/repository"
	"restaurant-app/utils"
)

type MenuItemService struct {
	RepoMenuItem repository.MenuItemRepository
}

func NewMenuItemService(repo repository.MenuItemRepository) *MenuItemService {
	return &MenuItemService{RepoMenuItem: repo}
}

func (m *MenuItemService) AddMenuService(menu *model.MenuItem) error {
	_, role, isAuthorized := utils.SessionRole()
	if !isAuthorized {
		return nil
	}
	if role != "Admin" {
		utils.SendJSONResponse(403, "Forbidden: only Admin can added menu", nil)

		return nil
	}
	err := utils.ReadBodyJSON("body.json", &menu)
	if err != nil {
		return nil
	}
	
	if menu.Name == "" {
		utils.SendJSONResponse(400, "name cannot empty", nil)
		return nil
	}
	if menu.Price <= 0 {
		utils.SendJSONResponse(400, "price cannot empty", nil)
		return nil
	}

	err = m.RepoMenuItem.Create(menu)
	if err != nil {
		utils.SendJSONResponse(500, err.Error(), nil)
		return nil
	}
	utils.SendJSONResponse(201, "successfully added menu", menu)
	return nil
}
