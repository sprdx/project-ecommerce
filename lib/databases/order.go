package databases

import (
	"project-ecommerce/config"
	"project-ecommerce/models"
)

func CreateOrder(newOrder *models.Order) (uint, error) {
	tx := config.DB.Create(&newOrder)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return newOrder.ID, nil
}

func CreateOrderDetails(newOrderDetail *models.OrderDetail) error {
	tx := config.DB.Create(&newOrderDetail)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func GetOrders(UserId int) (interface{}, error) {
	var order []models.Order
	tx := config.DB.Where("orders.user_id = ?", UserId).Find(&order)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return order, nil
}

func GetOrderDetails(orderId int) (interface{}, error) {
	var orderDetail []models.OrderDetail
	tx := config.DB.Where("order_details.order_id = ?", orderId).Find(&orderDetail)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return orderDetail, nil
}
