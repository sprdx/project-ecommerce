package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Product_name string  `json:"name" form:"name"`
	Category     string  `json:"category" form:"category"`
	Price        float64 `json:"price" form:"price"`
	Stock        uint    `json:"stock" form:"stock"`
	Detail       string  `json:"detail" form:"detail"`
	Rating       float64 `json:"rating" form:"rating"`
	Photo        string  `json:"photo" form:"photo"`
	UserID       uint
}

func (newProduct *Product) Validate() string {
	validate := validator.New()
	err := validate.Var(newProduct.Product_name, "required,min=3,max=35,startsnotwith= ,endsnotwith= ")
	if err != nil {
		return "Invalid name of product"
	}
	err = validate.Var(newProduct.Category, "required,alpha,startsnotwith= ,endsnotwith= ")
	if err != nil {
		return "Invalid product category"
	}
	err = validate.Var(newProduct.Price, "required,number,gte=1000")
	if err != nil {
		return "Invalid price"
	}
	err = validate.Var(newProduct.Stock, "required,number,gt=0")
	if err != nil {
		return "Invalid stock"
	}
	err = validate.Var(newProduct.Detail, "required")
	if err != nil {
		return "Invalid detail"
	}
	// err = validate.Var(newProduct.Rating, "required,number,gt=0,lte=5")
	// if err != nil {
	// 	return "Invalid rating"
	// }
	return "OK"
}
