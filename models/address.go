package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Address_detail string `json:"detail" form:"detail"`
	UserID         uint
}
