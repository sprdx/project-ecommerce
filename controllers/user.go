package controllers

import (
	"net/http"
	"strconv"

	"project-ecommerce/config"
	"project-ecommerce/lib/databases"
	"project-ecommerce/middlewares"
	"project-ecommerce/models"
	"project-ecommerce/responses"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserController(c echo.Context) error {
	var newUser models.User
	c.Bind(&newUser)

	// validate each fields that have be inputted by user
	message := newUser.Validate()
	if message != "OK" {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse(message))
	}
	var user models.User
	tx := config.DB.Where("email = ?", newUser.Email).First(&user)
	if tx.RowsAffected > 0 {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("Email has been registered"))
	}
	// Encrypt user's password before insert into database
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	newUser.Password = string(hashPassword)

	// Insert value of user struct into database
	_, err := databases.CreateUser(&newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("A database error occured"))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponseNonData("Congratulation! User created successfully"))
}

func LoginUserController(c echo.Context) error {
	var userData models.User
	c.Bind(&userData)

	// Check if data of user login is exist and correct in database
	token, err := databases.LoginUser(&userData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("Email or password is incorrect"))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponseData("Successful login", token))
}

func GetUserByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("Invalid ID"))
	}

	// Check if id from token is match to inputted id
	TokenUserId := middlewares.ExtractTokenUserId(c)
	if TokenUserId != id {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("Access forbidden"))
	}

	user, err := databases.GetUserById(TokenUserId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("A database error occured"))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponseData("Successful Operation", user))
}
