package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `json:"username" form:"username"`
	Email     string `gorm:"unique" json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Token     string
	Birthdate time.Time `gorm:"type:date"`
	Gender    string
	Phone     string
	Photo     string
	Products  []Product
	Carts     []Cart
	Orders    []Order
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
