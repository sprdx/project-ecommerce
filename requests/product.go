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

func (data *CreateProduct) ValidateUpdateProduct() string {
	validate := validator.New()
	if len(data.ProductName) != 0 {
		err := validate.Var(data.ProductName, "required,min=3,max=35,startsnotwith= ,endsnotwith= ")
		if err != nil {
			return "Invalid name of product"
		}
	}
	if len(data.Category) != 0 {
		err := validate.Var(data.Category, "required,alpha,startsnotwith= ,endsnotwith= ")
		if err != nil {
			return "Invalid product category"
		}
	}
	if data.Price != 0 {
		err := validate.Var(data.Price, "required,number,gte=1000")
		if err != nil {
			return "Invalid price"
		}
	}
	if data.Stock != 0 {
		err := validate.Var(data.Stock, "required,number,gt=0")
		if err != nil {
			return "Invalid stock"
		}
	}
	if len(data.Detail) != 0 {
		err := validate.Var(data.Detail, "required")
		if err != nil {
			return "Invalid detail"
		}
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

func BindUpdateProductData(c echo.Context, data *CreateProduct, prod *models.Product) string {
	c.Bind(&data)
	var countUpdate int
	message := data.ValidateUpdateProduct()
	if message != "VALID" {
		return message
	}

	if data.ProductName != "" {
		prod.ProductName = data.ProductName
		countUpdate++
	}
	if data.Category != "" {
		prod.Category = data.Category
		countUpdate++
	}
	if data.Price != 0 {
		prod.Price = data.Price
		countUpdate++
	}
	if data.Stock != 0 {
		prod.Stock = data.Stock
		countUpdate++
	}
	if data.Detail != "" {
		prod.Detail = data.Detail
		countUpdate++
	}
	if countUpdate == 0 {
		message = "No data have been updated"
	}

	return message
}
