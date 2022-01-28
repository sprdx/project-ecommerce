package controllers

import (
	"project-ecommerce/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateCartControllers(c echo.Context) error {
	id := strconv.Atoi(c.Param("id"))

	var cart models.Cart
	qty := c.Bind(&cart)
}
