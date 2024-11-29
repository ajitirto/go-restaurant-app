package database

import (
	"fmt"
	"gorm.io/gorm"

	"github.com/ajitirto/go-restaurant-app/internal/model"
	"github.com/ajitirto/go-restaurant-app/internal/model/constant"
)

func seedDB(db *gorm.DB) {
	db.AutoMigrate(&model.MenuItem{}, &model.Order{}, &model.ProductOrder{})
	
	foodMenu := []model.MenuItem {
		{
			Name:      "Bakmie",
			OrderCode: "bakmie",
			Price:     37500,
			Type:      constant.Food,
		},
		{
			Name:      "Ayam Rica-Rica",
			OrderCode: "ayam_rica_rica",
			Price:     41250,
			Type:      constant.Food,
		},
	}

	drinksMenu := []model.MenuItem {
		{
			Name:      "Es Teh",
			OrderCode: "es_teh",
			Price:     4000,
			Type:      constant.Drink,
		},
		{
			Name:      "Es Teh Manis",
			OrderCode: "es_teh_manis",
			Price:     5000,
			Type:      constant.Drink,
		},
		{
			Name:      "Air Mineral",
			OrderCode: "air_mineral",
			Price:     7000,
			Type:      constant.Drink,
		},
		{
			Name:      "Jus Apel",
			OrderCode: "jus_apel",
			Price:     14000,
			Type:      constant.Drink,
		},
	}
	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
	fmt.Println("Seeding db data ... ")
	db.Create(&foodMenu)
	db.Create(&drinksMenu)
	}
}