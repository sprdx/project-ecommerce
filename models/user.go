package models

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `json:"username" form:"username"`
	Email     string `gorm:"unique" json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Token     string
	Birthdate string
	Gender    string
	Phone     string
	Photo     string
	Product   []Product
	Cart      []Cart
}

type GetUser struct {
	ID        uint
	Name      string
	Email     string
	Birthdate string
	Gender    string
	Phone     string
	Photo     string
}

func (newUser *User) Validate() string {
	validate := validator.New()
	err := validate.Var(newUser.Username, "required,min=3,max=35,startsnotwith= ,endsnotwith= ")
	if err != nil {
		return "Invalid name"
	}
	if !regexp.MustCompile(`^[a-zA-Z]+(\s{1}[a-zA-Z]+)*$`).MatchString(newUser.Username) || strings.Count(newUser.Username, " ") > 2 {
		return "Allowed full name is only alphabet and contain maximum 3 words"
	}
	err = validate.Var(newUser.Email, "required,email")
	if err != nil {
		return "Invalid email"
	}
	err = validate.Var(newUser.Password, "required,min=8")
	if err != nil {
		return "Invalid password"
	}
	return "OK"
}
