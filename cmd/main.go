package main

import (
	"github.com/labstack/echo/v4"

	mRepo "github.com/ajitirto/go-restaurant-app/internal/repository/menu"
	oRepo "github.com/ajitirto/go-restaurant-app/internal/repository/order"	
	rUsecase "github.com/ajitirto/go-restaurant-app/internal/usecase/resto"

	"github.com/ajitirto/go-restaurant-app/internal/database"
	"github.com/ajitirto/go-restaurant-app/internal/delivery/rest"
)

func main() {

	e := echo.New()
	db := database.GetDB("host=172.21.0.2 port=5432 user=postgres password=postgres dbname=go_resto sslmode=disable")

	menuRepo := mRepo.GetRepository(db)
	orderRepo := oRepo.GetRepository(db)

	restoUsecase := rUsecase.GetUsecase(menuRepo, orderRepo)

	h := rest.NewHandler(restoUsecase)

	rest.LoadRoutes(e, h)

	e.Logger.Fatal(e.Start((":14045")))
}
