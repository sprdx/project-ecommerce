package controllers

import (
	"fmt"
	"net/http"
	"project-ecommerce/lib/databases"
	"project-ecommerce/middlewares"
	"project-ecommerce/requests"
	"project-ecommerce/responses"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateProductController(c echo.Context) error {
	var data requests.CreateProduct

	product, message := requests.BindProductData(c, &data)
	if message != "VALID" {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse(message))
	}

	_, err := databases.CreateProduct(product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("A database error ocurred"))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponseNonData("Successful creating a new product"))
}

func GetAllProductsController(c echo.Context) error {
	products, err := databases.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("A database error ocurred"))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponseData("Successful operation", products))
}

func GetProductByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("Invalid ID"))
	}

	product, err := databases.GetProductById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("A database error occured"))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponseData("Successful Operation", product))
}

func UpdateProductController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("Invalid ID"))
	}

	product, err := databases.GetTheProduct(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("Data not found"))
	}
	fmt.Println(product)
	// Check if id from token is match to inputted id
	tokenUserId := middlewares.ExtractTokenUserId(c)
	fmt.Println(tokenUserId, "==", int(product.UserID))
	if tokenUserId != int(product.UserID) {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("Access forbidden"))
	}

	var updateData requests.CreateProduct

	message := requests.BindUpdateProductData(c, &updateData, &product)
	if message != "VALID" {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse(message))
	}

	err = databases.UpdateProduct(&product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("A database error occured"))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponseNonData("Successful Operation"))
}

func DeleteProductController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("Invalid ID"))
	}

	product, err := databases.GetTheProduct(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("Data not found"))
	}

	// Check if id from token is match to inputted id
	tokenUserId := middlewares.ExtractTokenUserId(c)
	if tokenUserId != int(product.UserID) {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("Access forbidden"))
	}

	err = databases.DeleteProduct(&product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("A database error occured"))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponseNonData("Successful Operation"))
}
