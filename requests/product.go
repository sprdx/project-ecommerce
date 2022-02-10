package requests

import (
	"project-ecommerce/middlewares"
	"project-ecommerce/models"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CreateProduct struct {
	ProductName string  `json:"product_name" form:"product_name"`
	Category    string  `json:"category" form:"category"`
	Price       float64 `json:"price" form:"price"`
	Stock       uint    `json:"stock" form:"stock"`
	Detail      string  `json:"detail" form:"detail"`
}

func (data *CreateProduct) Validate() string {
	validate := validator.New()
	err := validate.Var(data.ProductName, "required,min=3,max=35,startsnotwith= ,endsnotwith= ")
	if err != nil {
		return "Invalid name of product"
	}
	err = validate.Var(data.Category, "required,alpha,startsnotwith= ,endsnotwith= ")
	if err != nil {
		return "Invalid product category"
	}
	err = validate.Var(data.Price, "required,number,gte=1000")
	if err != nil {
		return "Invalid price"
	}
	err = validate.Var(data.Stock, "required,number,gt=0")
	if err != nil {
		return "Invalid stock"
	}
	err = validate.Var(data.Detail, "required")
	if err != nil {
		return "Invalid detail"
	}
	return "VALID"
}

func BindProductData(c echo.Context, data *CreateProduct) (*models.Product, string) {
	c.Bind(&data)
	var product models.Product
	message := data.Validate()
	if message != "VALID" {
		return &product, message
	}

	product.ProductName = data.ProductName
	product.Category = data.Category
	product.Price = data.Price
	product.Stock = data.Stock
	product.Detail = data.Detail
	product.UserID = uint(middlewares.ExtractTokenUserId(c))

	return &product, message
}
