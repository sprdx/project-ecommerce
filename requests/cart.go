package requests

import (
	"project-ecommerce/lib/databases"
	"project-ecommerce/middlewares"
	"project-ecommerce/models"

	"github.com/labstack/echo/v4"
)

type CreateCart struct {
	ProductID uint `json:"product_id" form:"product_id"`
	Quantity  uint `json:"quantity" form:"quantity"`
}

func (data *CreateCart) Validate(c echo.Context) (uint, string) {
	prod, err := databases.GetTheProduct(int(data.ProductID))
	if err != nil {
		return 0, "Product not found"
	}
	userId := uint(middlewares.ExtractTokenUserId(c))
	if prod.UserID == userId {
		return userId, "Access forbidden. Seller not allowed to buy her/his stuff"
	}
	if data.Quantity < 1 {
		return userId, "Minimum order quantity is 1"
	}
	if data.Quantity > prod.Stock {
		return userId, "Order quantity should not greater than stock of the product"
	}
	return userId, "VALID"
}

func BindCartData(c echo.Context, data *CreateCart) (*models.Cart, string) {
	c.Bind(data)
	var cart models.Cart
	userId, message := data.Validate(c)
	if message != "VALID" {
		return &cart, message
	}
	cart.ProductID = data.ProductID
	cart.Quantity = data.Quantity
	cart.UserID = userId

	return &cart, "VALID"
}
