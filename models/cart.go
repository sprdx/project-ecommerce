package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Quantity    uint    `json:"quantity" form:"quantity"`
	TotalPrice  float64 `json:"total_price" form:"total_price"`
	ProductID   uint
	UserID      uint
	OrderDetail OrderDetail
}

type GetCart struct {
	ID          uint
	ProductName string
	Quantity    uint
	TotalPrice  float64
}
