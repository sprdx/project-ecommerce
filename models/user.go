package models

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	User_name string `json:"name" form:"name"`
	Email     string `gorm:"unique" json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Token     string
	Birthdate string
	Gender    string
	Phone     string
	Photo     string
	Product   []Product
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
	err := validate.Var(newUser.User_name, "required,min=3,max=35,startsnotwith= ,endsnotwith= ")
	if err != nil {
		return "Invalid name"
	}
	if !regexp.MustCompile(`^[a-zA-Z]+(\s{1}[a-zA-Z]+)*$`).MatchString(newUser.User_name) || strings.Count(newUser.User_name, " ") > 2 {
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
