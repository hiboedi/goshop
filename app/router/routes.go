package router

import (
	"github.com/hiboedi/zenshop/app/controllers"
	"github.com/hiboedi/zenshop/app/exception"
	"github.com/julienschmidt/httprouter"
)

func RouterInit(productController controllers.ProductController, userController controllers.UserController, orderController controllers.OrderController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/products", productController.FindAll)
	router.GET("/api/products/:productId", productController.FindById)
	router.POST("/api/products", productController.Create)
	router.PUT("/api/products/:productId", productController.Update)
	router.DELETE("/api/products/:productId", productController.Delete)

	router.POST("/api/users/login", userController.Login)
	router.POST("/api/users", userController.Create)
	router.PUT("/api/users/:userId", userController.Update)

	router.POST("/api/orders", orderController.Create)
	router.PUT("/api/orders/:orderId", orderController.Update)
	router.DELETE("/api/orders/:orderId", orderController.Delete)
	router.GET("/api/orders/:orderId", orderController.FindById)
	router.GET("/api/orders", orderController.FindAllOrByPaymentStatus)

	router.PanicHandler = exception.ErrorHandler

	return router
}
