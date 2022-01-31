package controllers

import (
	"fmt"
	"net/http"
	"project-ecommerce/lib/databases"
	"project-ecommerce/models"
	"project-ecommerce/responses"

	"github.com/labstack/echo/v4"
)

func CreateOrderControllers(c echo.Context) error {

	fmt.Println("Hello")
	var newPayment models.Payment
	c.Bind(&newPayment)
	fmt.Println(newPayment)
	err := databases.CreateOrder(newPayment)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("A database error ocurred"))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponseNonData("Successful Operation bang"))
}

func HelloControllers(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World!")
}
