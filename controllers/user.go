package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"project-ecommerce/lib/databases"
	"project-ecommerce/middlewares"
	"project-ecommerce/models"
	"project-ecommerce/requests"
	"project-ecommerce/responses"

	"github.com/labstack/echo/v4"
)

func CreateUserController(c echo.Context) error {
	var data requests.CreateUser
	var newUser models.User
	fmt.Println("user", newUser)
	message := requests.BindUserData(c, &data, &newUser)
	if message != "VALID" {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse(message))
	}
	fmt.Println("user", newUser)
	// Insert value of user struct into database
	_, err := databases.CreateUser(&newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("A database error occured"))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponseNonData("Congratulation! User created successfully"))
}

func LoginUserController(c echo.Context) error {
	var loginData requests.LoginUser

	user, message := requests.BindLoginData(c, &loginData)
	if message != "VALID" {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse(message))
	}

	data, err := databases.LoginUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("A database error occured"))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponseData("Successful login", data))
}

func GetUserByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("Invalid ID"))
	}

	// Check if id from token is match to inputted id
	tokenUserId := middlewares.ExtractTokenUserId(c)
	if tokenUserId != id {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("Access forbidden"))
	}

	data, err := databases.GetUserById(tokenUserId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("A database error occured"))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponseData("Successful Operation", data))
}

func UpdateUserByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("Invalid ID"))
	}

	// Check if id from token is match to inputted id
	tokenUserId := middlewares.ExtractTokenUserId(c)
	if tokenUserId != id {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("Access forbidden"))
	}
	var updateData requests.UpdateUser

	user, message := requests.BindUpdateData(c, &updateData, tokenUserId)
	if message != "VALID" {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse(message))
	}

	err = databases.UpdateUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("A database error occured"))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponseNonData("Successful Operation"))
}
