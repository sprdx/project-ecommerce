package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName string  `json:"product_name" form:"product_name"`
	Category    string  `json:"category" form:"category"`
	Price       float64 `json:"price" form:"price"`
	Stock       uint    `json:"stock" form:"stock"`
	Detail      string  `json:"detail" form:"detail"`
	Rating      float64 `json:"rating" form:"rating"`
	Photo       string  `json:"photo" form:"photo"`
	UserID      uint
	Carts       []Cart
}

type GetProduct struct {
	ProductName string
	Category    string
	Price       float64
	Stock       uint
	Detail      string
	Rating      float64
	Photo       string
	Seller      string
}
