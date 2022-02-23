package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Quantity     uint
	TotalPrice   float64 `json:"total_price" form:"total_price"`
	Status       string  `gorm:"default:on processing"`
	UserID       uint
	OrderDetails []OrderDetail
}

type OrderDetail struct {
	gorm.Model
	CartID  uint
	OrderID uint
}

type GetOrder struct {
	ID         uint
	CreatedAt  string
	Quantity   uint
	TotalPrice float64
	Status     string
	Buyer      string
}

type GetOrderDetails struct {
	ID         uint
	CreatedAt  string
	Quantity   uint
	TotalPrice float64
	OrderID    uint
}
