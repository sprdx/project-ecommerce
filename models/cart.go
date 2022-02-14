package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Quantity    uint `json:"quantity" form:"quantity"`
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
