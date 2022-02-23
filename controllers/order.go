package controllers

import (
	"net/http"
	"project-ecommerce/lib/databases"
	"project-ecommerce/middlewares"
	"project-ecommerce/models"
	"project-ecommerce/requests"
	"project-ecommerce/responses"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateOrderController(c echo.Context) error {
	var newOrder requests.CreateOrder
	order, message := requests.BindOrderData(c, &newOrder)
	if message != "VALID" {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse(message))
	}

	orderId, err := databases.CreateOrder(order)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("A database error when create order"))
	}

	for _, cartId := range newOrder.CartIDList {
		var newOrderDetail models.OrderDetail
		newOrderDetail.OrderID = orderId
		newOrderDetail.CartID = cartId
		err := databases.CreateOrderDetails(&newOrderDetail)
		if err != nil {
			return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("A database error when create order details"))
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
	orderId, _ := strconv.Atoi(c.Param("id"))
	data, err := databases.GetOrderDetails(orderId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("A database error ocurred"))
	}
	return c.JSON(http.StatusOK, responses.SuccessResponseData("Successful Operation bang", data))
}
