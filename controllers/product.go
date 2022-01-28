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

func CreateProductController(c echo.Context) error {
	var newProduct models.Product
	c.Bind(&newProduct)
	newProduct.UserID = uint(middlewares.ExtractTokenUserId(c))

	message := newProduct.Validate()
	if message != "OK" {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse(message))
	}

	_, err := databases.CreateProduct(&newProduct)
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
