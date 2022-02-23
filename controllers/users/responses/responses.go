package responses

import (
	"project-ecommerce/bussinesses/users"
	"time"
)

type UserResponse struct {
	ID          uint
	Username    string
	Email       string
	Birthdate   time.Time
	Gender      string
	PhoneNumber string
}

func FromDomain(domain users.Domain) UserResponse {
	return UserResponse{
		ID:          domain.Id,
		Username:    domain.Username,
		Email:       domain.Email,
		Birthdate:   domain.Birthdate,
		Gender:      domain.Gender,
		PhoneNumber: domain.PhoneNumber,
	}
}
