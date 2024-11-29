package menu

import (
	"gorm.io/gorm"
	
	"github.com/ajitirto/go-restaurant-app/internal/model"
	"github.com/ajitirto/go-restaurant-app/internal/model/constant"
)

type menuRepo struct {
	db *gorm.DB
}

func GetRepository(db *gorm.DB) *menuRepo {
	return &menuRepo{
		db: db,
	}
}

func (m *menuRepo) GetMenuList(menuType string) ([]model.MenuItem, error) {
	menuData := make([]model.MenuItem, 0)

	if err := m.db.Where(model.MenuItem{Type: constant.MenuType(menuType)}).Find(&menuData).Error; err != nil {
		return nil, err
	}

	return menuData, nil
}

func (m *menuRepo) GetMenu(OrderCode string) (model.MenuItem, error) {
	var menuData model.MenuItem

	if err := m.db.Where(model.MenuItem{OrderCode: OrderCode}).First(&menuData).Error; err != nil {
		return model.MenuItem{}, err
	}

	return menuData, nil
}