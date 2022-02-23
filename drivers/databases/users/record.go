package users

import (
	"project-ecommerce/bussinesses/users"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id          uint
	Username    string
	Email       string
	Password    string
	Token       string
	Birthdate   time.Time
	Gender      string
	PhoneNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	// Products  []Product
	// Carts     []Cart
	// Orders    []Order
}

func (Users) TableName() string {
	return "users"
}

func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}

func (u *Users) ToDomain() users.Domain {
	return users.Domain{
		Id:          u.Id,
		Username:    u.Username,
		Email:       u.Email,
		Birthdate:   u.Birthdate,
		Gender:      u.Gender,
		PhoneNumber: u.PhoneNumber,
	}
}

func FromDomain(domain users.Domain) Users {
	return Users{
		Username:    domain.Username,
		Email:       domain.Email,
		Password:    domain.Password,
		Token:       domain.Token,
		Birthdate:   domain.Birthdate,
		Gender:      domain.Gender,
		PhoneNumber: domain.PhoneNumber,
	}
}

func ToListDomain(data []Users) []users.Domain {
	var listDomain []users.Domain
	for _, d := range data {
		listDomain = append(listDomain, d.ToDomain())
	}
	return listDomain
}
