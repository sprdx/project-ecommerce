package models

import "gorm.io/gorm"

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

type OrderRequest struct {
	CartIDList []uint `json:"cart_id_list" form:"cart_id_list"`
}
