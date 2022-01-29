package controllers

import (
	"net/http"
	"project-ecommerce/lib/databases"
	"project-ecommerce/middlewares"
	"project-ecommerce/models"
	"project-ecommerce/responses"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateCartControllers(c echo.Context) error {
	var newCart models.Cart
	id, _ := strconv.Atoi(c.Param("id"))
	newCart.ProductID = uint(id)
	newCart.UserID = uint(middlewares.ExtractTokenUserId(c))
	c.Bind(&newCart)

	_, err := databases.CreateCart(&newCart)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("A database error ocurred"))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponseNonData("Successful Operation"))
}

func GetCartControllers(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)
	data, err := databases.GetCart(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("A database error ocurred"))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponseData("Successful Operation", data))
}

func DeleteCartControllers(c echo.Context) error {
	productId, _ := strconv.Atoi(c.Param("id"))
	userId := middlewares.ExtractTokenUserId(c)
	err := databases.DeleteCart(userId, productId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("A database error ocurred"))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponseNonData("Successful Operation"))
}
