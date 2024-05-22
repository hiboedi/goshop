package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/hiboedi/zenshop/app/controllers"
	"github.com/hiboedi/zenshop/app/database"
	"github.com/hiboedi/zenshop/app/helpers"
	"github.com/hiboedi/zenshop/app/repository"
	"github.com/hiboedi/zenshop/app/router"
	"github.com/hiboedi/zenshop/app/services"
)

func main() {
	db := database.DbInit()
	validate := validator.New()

	productRepo := repository.NewProductRepo()
	productService := services.NewProductService(productRepo, db, validate)
	productController := controllers.NewProductController(productService)

	userRepo := repository.NewUserRepo()
	userService := services.NewUserService(userRepo, db, validate)
	userController := controllers.NewUserController(userService)

	orderRepo := repository.NewOrderRepo()
	orderService := services.NewOrderService(orderRepo, db, validate)
	orderController := controllers.NewOrderController(orderService)

	router := router.RouterInit(productController, userController, orderController)
	database.Migrate()

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: router,
	}
	fmt.Println("starting on port :8000")

	err := server.ListenAndServe()
	helpers.PanicIfError(err)
}
