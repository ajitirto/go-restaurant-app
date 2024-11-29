package menu

import (
	"github.com/ajitirto/go-restaurant-app/internal/model"
)

type Repository interface {
	GetMenuList(menuType string) ([]model.MenuItem, error)
	GetMenu(OrderCode string) (model.MenuItem, error)
}
