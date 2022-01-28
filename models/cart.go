package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Quantity  uint `json:"qty" form:"qty"`
	ProductID uint
	UserID    uint
}
