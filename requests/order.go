package requests

import (
	"fmt"
	"project-ecommerce/config"
	"project-ecommerce/lib/databases"
	"project-ecommerce/middlewares"
	"project-ecommerce/models"

	"github.com/labstack/echo/v4"
)

type CreateOrder struct {
	CartIDList []uint `json:"cart_id_list" form:"cart_id_list"`
}

func BindOrderData(c echo.Context, data *CreateOrder) (*models.Order, string) {
	c.Bind(data)
	var cart models.Cart
	var order models.Order
	order.UserID = uint(middlewares.ExtractTokenUserId(c))

	check := config.DB.Where("user_id = ?", order.UserID).Find(&cart, data.CartIDList)
	if check.RowsAffected == 0 {
		return &order, "Cart ID is invalid"
	}

	for _, ID := range data.CartIDList {
		Quantity, TotalPrice, _ := databases.GetDetailCart(int(ID))
		fmt.Println("Qty", Quantity)
		fmt.Println("Total", TotalPrice)
		order.Quantity += Quantity
		order.TotalPrice += TotalPrice
	}

	return &order, "VALID"
}
