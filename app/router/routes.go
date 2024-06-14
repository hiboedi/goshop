package router

import (
	"github.com/gorilla/mux"
	"github.com/hiboedi/zenshop/app/controllers"
	"github.com/hiboedi/zenshop/app/middleware"
)

func RouterInit(productController controllers.ProductController, userController controllers.UserController, orderController controllers.OrderController, addressController controllers.AddressController) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/users/login", userController.Login).Methods("POST")
	router.HandleFunc("/api/users", userController.Create).Methods("POST")
	router.HandleFunc("/api/users/{userId}", userController.Update).Methods("PUT")

	router.HandleFunc("/api/users/addresses", addressController.Create).Methods("POST")
	router.HandleFunc("/api/users/addresses", addressController.FindByUserId).Methods("GET")
	router.HandleFunc("/api/users/addresses/{addressId}", addressController.Delete).Methods("DELETE")
	router.HandleFunc("/api/users/addresses/{addressId}", addressController.Update).Methods("PUT")
	router.HandleFunc("/api/users/addresses/{addressId}", addressController.FindById).Methods("GET")

	router.HandleFunc("/api/products", productController.FindAll).Methods("GET")
	router.HandleFunc("/api/products/{productId}", productController.FindById).Methods("GET")
	router.HandleFunc("/api/products", productController.Create).Methods("POST")
	router.HandleFunc("/api/products/{productId}", productController.Update).Methods("PUT")
	router.HandleFunc("/api/products/{productId}", productController.Delete).Methods("DELETE")

	router.HandleFunc("/api/orders", orderController.Create).Methods("POST")
	router.HandleFunc("/api/orders/{orderId}", orderController.Update).Methods("PUT")
	router.HandleFunc("/api/orders/{orderId}", orderController.Delete).Methods("DELETE")
	router.HandleFunc("/api/orders/{orderId}", orderController.FindById).Methods("GET")
	router.HandleFunc("/api/orders", orderController.FindAllOrByPaymentStatus).Methods("GET")

	// Error handling middleware can be added as needed

	router.Use(middleware.RecoverMiddleware)

	return router
}
