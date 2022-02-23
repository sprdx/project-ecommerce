package users

import (
	"net/http"
	"project-ecommerce/bussinesses/users"
	"project-ecommerce/controllers/users/requests"
	"project-ecommerce/controllers/users/responses"
	res "project-ecommerce/responses"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Usecase users.Usecase
}

func NewUserController(u users.Usecase) *UserController {
	return &UserController{
		Usecase: u,
	}
}

func (u *UserController) Register(c echo.Context) error {
	userRegister := requests.RegisterRequest{}
	c.Bind(&userRegister)

	// check if email is already registered
	ctx := c.Request().Context()
	// user, _ := a.Usecase.FindByEmail(ctx, userRegister.Email)
	// if user.Id != 0 {
	// 	return controllers.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrUserAlreadyExists)
	// }

	userDomain := users.Domain{
		Username: userRegister.Username,
		Email:    userRegister.Email,
		Password: userRegister.Password,
	}

	user, err := u.Usecase.Register(ctx, userDomain)
	if user.Id == 0 {
		return c.JSON(http.StatusBadRequest, res.BadRequestResponse("error bad request"))
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, res.BadRequestResponse("A database error occured"))
	}

	registerResponse := responses.UserResponse{
		ID:          user.Id,
		Username:    user.Username,
		Email:       user.Email,
		Birthdate:   user.Birthdate,
		PhoneNumber: user.PhoneNumber,
		Gender:      user.Gender,
	}
	return c.JSON(http.StatusOK, res.SuccessResponseData("Congratulation! User created successfully", registerResponse))
}
