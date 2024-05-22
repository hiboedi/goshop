package test

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/hiboedi/zenshop/app/controllers"
	"github.com/hiboedi/zenshop/app/helpers"
	"github.com/hiboedi/zenshop/app/middleware"
	"github.com/hiboedi/zenshop/app/models"
	"github.com/hiboedi/zenshop/app/repository"
	"github.com/hiboedi/zenshop/app/router"
	"github.com/hiboedi/zenshop/app/services"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBTestSetup() *gorm.DB {
	dsn := "root:@tcp(localhost:3306)/langgo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helpers.PanicIfError(err)
	return db
}

func Migrate() {
	db := DBTestSetup()
	err := db.AutoMigrate(
		&models.Product{},
		&models.User{},
		&models.Order{},
	)
	helpers.PanicIfError(err)

	fmt.Println("Migration success")
}

func SetUpRouter() http.Handler {
	db := DBTestSetup()
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
	Migrate()

	return middleware.NewAuthMiddleware(router)
}
