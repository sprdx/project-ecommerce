package controllers

import (
	"fmt"
	"net/http"
	"project-ecommerce/lib/databases"
	"project-ecommerce/middlewares"
	"project-ecommerce/models"
	"project-ecommerce/responses"

	"github.com/labstack/echo/v4"
)

func CreateOrderControllers(c echo.Context) error {
	var newOrderRequest models.OrderRequest
	var newOrder models.Order
	c.Bind(&newOrderRequest)

	for _, ID := range newOrderRequest.CartIDList {
		Quantity, TotalPrice, _ := databases.GetDetailCart(int(ID))
		fmt.Println("Qty", Quantity)
		fmt.Println("Total", TotalPrice)
		newOrder.Quantity += Quantity
		newOrder.TotalPrice += TotalPrice
	}
	newOrder.UserID = uint(middlewares.ExtractTokenUserId(c))

	err := databases.CreateOrder(&newOrder)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("A database error ocurred"))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponseNonData("Successful Operation bang"))
}
