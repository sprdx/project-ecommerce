package routes

import (
	"project-ecommerce/controllers/users"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserController *users.UserController
}

func (c ControllerList) InitRoutes(e *echo.Echo) {
	e.Debug = true

	// e.POST("/login", controllers.LoginUserController)
	e.POST("/users", c.UserController.Register)
	// e.GET("/products", controllers.GetAllProductsController)
	// e.GET("/products/:id", controllers.GetProductByIdController)

	// r := e.Group("/jwt")
	// r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	// // r.GET("/users/:id", controllers.GetUserByIdController)
	// // r.PATCH("/users/:id", controllers.UpdateUserByIdController)
	// r.POST("/products", controllers.CreateProductController)
	// r.PUT("/products/:id", controllers.UpdateProductController)
	// r.DELETE("/products/:id", controllers.DeleteProductController)
	// r.POST("/carts", controllers.CreateCartController)
	// r.GET("/carts", controllers.GetUserCartsController)
	// r.DELETE("/carts/products/:id", controllers.DeleteCartController)
	// r.POST("/orders", controllers.CreateOrderController)
	// r.GET("/orders", controllers.GetUserOrdersController)
	// r.GET("/orders/:id", controllers.GetUserOrderDetailsController)
}
