package requests

import (
	"project-ecommerce/lib/databases"
	"project-ecommerce/middlewares"
	"project-ecommerce/models"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type CreateUser struct {
	Username string `json:"username" form:"username"`
	Email    string `gorm:"unique" json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type LoginUser struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UpdateUser struct {
	Username  string `json:"username" form:"username"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Birthdate string `json:"birthdate" form:"birthdate"`
	Gender    string `json:"gender" form:"gender"`
	Phone     string `json:"phone" form:"phone"`
	Photo     string `json:"photo" form:"photo"`
}

func (data *CreateUser) Validate() string {
	validate := validator.New()
	err := validate.Var(data.Username, "required,min=3,max=35,startsnotwith= ,endsnotwith= ")
	if err != nil {
		return "Invalid name"
	}
	if !regexp.MustCompile(`^[a-zA-Z]+(\s{1}[a-zA-Z]+)*$`).MatchString(data.Username) || strings.Count(data.Username, " ") > 2 {
		return "Allowed full name is only alphabet and contain maximum 3 words"
	}
	err = validate.Var(data.Email, "required,email")
	if err != nil {
		return "Invalid email"
	}
	_, count := databases.GetUserByEmail(data.Email)
	if count != 0 {
		return "Email has been registered"
	}
	err = validate.Var(data.Password, "required,min=8")
	if err != nil {
		return "Invalid password"
	}
	return "VALID"
}

func (data *UpdateUser) ValidateUpdateData() string {
	validate := validator.New()
	if len(data.Username) != 0 {
		err := validate.Var(data.Username, "required,min=3,max=35,startsnotwith= ,endsnotwith= ")
		if err != nil {
			return "Invalid name"
		}
		if !regexp.MustCompile(`^[a-zA-Z]+(\s{1}[a-zA-Z]+)*$`).MatchString(data.Username) || strings.Count(data.Username, " ") > 2 {
			return "Allowed full name is only alphabet and contain maximum 3 words"
		}
	}
	if len(data.Email) != 0 {
		err := validate.Var(data.Email, "required,email")
		if err != nil {
			return "Invalid email"
		}
		_, count := databases.GetUserByEmail(data.Email)
		if count != 0 {
			return "Email has been registered"
		}
	}
	if len(data.Password) != 0 {
		err := validate.Var(data.Password, "required,min=8")
		if err != nil {
			return "Invalid password"
		}
	}
	if len(data.Birthdate) != 0 {
		if regexp.MustCompile(`^(19[5-9][0-9])|(200[0-7])-(02-(0[1-9]|[12][0-9])|(0[469]|11)-(0[1-9]|[12][0-9]|30)|(0[13578]|1[02])-(0[1-9]|[12][0-9]|3[01]))$`).MatchString(data.Birthdate) {
			if year, _ := strconv.Atoi(data.Birthdate[:4]); year%4 != 0 && data.Birthdate[5:7] == "02" && data.Birthdate[8:] == "29" {
				return "Invalid birthdate. Example format valid 1999-01-31"
			}
		}
	}
	if len(data.Gender) != 0 {
		if data.Gender != "Male" {
			if data.Gender != "Female" {
				return "Invalid gender. Choose between Male or Female."
			}
		}
	}
	if len(data.Phone) != 0 {
		if !regexp.MustCompile(`^08[0-9]{9,11}$`).MatchString(data.Phone) {
			return "Invalid phone number. Example format is 08123456789"
		}
	}
	return "VALID"
}

func BindUserData(c echo.Context, data *CreateUser, user *models.User) string {
	c.Bind(&data)
	message := data.Validate()

	user.Username = data.Username
	user.Email = data.Email

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	user.Password = string(hashPassword)

	return message
}

func BindLoginData(c echo.Context, data *LoginUser) (*models.User, string) {
	c.Bind(&data)

	// Check if user's login email is exist in database
	user, count := databases.GetUserByEmail(data.Email)
	if count == 0 {
		return user, "Email or password is wrong"
	}

	// Check if inputed password is match to password in database
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		return user, "Email or password is wrong"
	}

	// Generate token by using user's ID
	user.Token, _ = middlewares.CreateToken(int(user.ID))

	return user, "VALID"
}

func BindUpdateData(c echo.Context, data *UpdateUser, id int) (*models.User, string) {
	c.Bind(&data)
	message := data.ValidateUpdateData()
	if message != "VALID" {
		var user *models.User
		return user, message
	}

	user, count := databases.GetUser(id)
	if count == 0 {
		return user, "User is not found"
	}

	if data.Username != "" {
		user.Username = data.Username
	}
	if data.Email != "" {
		user.Email = data.Email
	}
	if data.Birthdate != "" {
		user.Birthdate, _ = time.Parse("2006-01-02", data.Birthdate)
	}
	if data.Gender != "" {
		user.Gender = data.Gender
	}
	if data.Phone != "" {
		user.Phone = data.Phone
	}
	return user, message
}
