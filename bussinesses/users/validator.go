package users

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

func (data *Domain) ValidateNewUserData() string {
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
	// _, count := databases.GetUserByEmail(data.Email)
	// if count != 0 {
	// 	return "Email has been registered"
	// }
	err = validate.Var(data.Password, "required,min=8")
	if err != nil {
		return "Invalid password"
	}
	return "VALID"
}
