package routes

import (
	"project-ecommerce/constants"
	"project-ecommerce/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/login", controllers.LoginUserController)
	e.POST("/users", controllers.CreateUserController)
	e.GET("/products", controllers.GetAllProductsController)
	e.GET("/products/:id", controllers.GetProductByIdController)

	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	r.GET("/users/:id", controllers.GetUserByIdController)
	r.PATCH("/users/:id", controllers.UpdateUserByIdController)
	r.POST("/products", controllers.CreateProductController)
	r.PUT("/products/:id", controllers.UpdateProductController)
	r.POST("/carts/products/:id", controllers.CreateCartController)
	r.GET("/carts", controllers.GetUserCartsController)
	r.DELETE("/carts/products/:id", controllers.DeleteCartController)
	r.POST("/orders", controllers.CreateOrderController)
	r.GET("/orders", controllers.GetUserOrdersController)
	r.GET("/orders/:id", controllers.GetUserOrderDetailsController)

	return e
}
