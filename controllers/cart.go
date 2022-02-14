package controllers

import (
	"net/http"
	"project-ecommerce/lib/databases"
	"project-ecommerce/middlewares"
	"project-ecommerce/requests"
	"project-ecommerce/responses"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateCartController(c echo.Context) error {
	var data requests.CreateCart
	cart, message := requests.BindCartData(c, &data)
	if message != "VALID" {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse(message))
	}

	_, err := databases.CreateCart(cart)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("A database error ocurred"))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponseNonData("Successful Operation"))
}

func GetUserCartsController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	data, err := databases.GetCart(userId)
	if data == nil || err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("Cart is empty"))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponseData("Successful Operation", data))
}

func DeleteCartController(c echo.Context) error {
	productId, _ := strconv.Atoi(c.Param("id"))
	userId := middlewares.ExtractTokenUserId(c)
	err := databases.DeleteCart(userId, productId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("Cart ID is not found"))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponseNonData("Successful Operation"))
}
