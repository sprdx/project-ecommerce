package controllers

import (
	"fmt"
	"net/http"
	"project-ecommerce/lib/databases"
	"project-ecommerce/middlewares"
	"project-ecommerce/models"
	"project-ecommerce/responses"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateOrderController(c echo.Context) error {
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

	orderId, err := databases.CreateOrder(&newOrder)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("A database error ocurred"))
	}

	for _, ID := range newOrderRequest.CartIDList {
		var newOrderDetail models.OrderDetail
		newOrderDetail.OrderID = orderId
		cartId := ID
		newOrderDetail.CartID = cartId
		err := databases.CreateOrderDetails(&newOrderDetail)
		if err != nil {
			return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("A database error ocurred"))
		}
	}

	return c.JSON(http.StatusOK, responses.SuccessResponseNonData("Successful Operation bang"))
}

func GetUserOrdersController(c echo.Context) error {
	UserId := middlewares.ExtractTokenUserId(c)
	data, err := databases.GetOrders(UserId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("A database error ocurred"))
	}
	return c.JSON(http.StatusOK, responses.SuccessResponseData("Successful Operation bang", data))
}

func GetUserOrderDetailsController(c echo.Context) error {
	orderId, err := strconv.Atoi(c.Param("id"))
	data, err := databases.GetOrderDetails(orderId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("A database error ocurred"))
	}
	return c.JSON(http.StatusOK, responses.SuccessResponseData("Successful Operation bang", data))
}
