package routes

import (
	"project-ecommerce/constants"
	"project-ecommerce/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/users", controllers.CreateUserController)
	e.POST("/login", controllers.LoginUserController)

	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	r.GET("/users/:id", controllers.GetUserByIdController)
	r.POST("/products", controllers.CreateProductController)
	return e
}
