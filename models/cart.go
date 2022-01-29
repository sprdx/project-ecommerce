package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Quantity    uint    `json:"qty" form:"qty"`
	Total_Price float64 `json:"total" form:"total"`
	ProductID   uint
	UserID      uint
}

type GetCart struct {
	Product_Name string
	Quantity     uint
	Total_Price  float64
}
