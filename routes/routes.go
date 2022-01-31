package routes

import (
	"project-ecommerce/constants"
	"project-ecommerce/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/ping", controllers.HelloControllers)
	e.POST("/login", controllers.LoginUserController)
	e.POST("/users", controllers.CreateUserController)
	e.GET("/products", controllers.GetAllProductsController)
	e.GET("/products/:id", controllers.GetProductByIdController)
	e.POST("/orders", controllers.CreateOrderControllers)

	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	r.GET("/users/:id", controllers.GetUserByIdController)
	r.POST("/products", controllers.CreateProductController)
	r.POST("/carts/products/:id", controllers.CreateCartControllers)
	r.GET("/carts", controllers.GetCartControllers)
	r.DELETE("/carts/products/:id", controllers.DeleteCartControllers)

	return e
}
