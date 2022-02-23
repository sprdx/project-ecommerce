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
	var cart models.Cart
	var product models.Product

	getCart := config.DB.Where("id = ?", newOrderDetail.CartID).Find(&cart)
	if getCart.Error != nil {
		return getCart.Error
	}
	getProduct := config.DB.Where("id = ?", cart.ProductID).Find(&product)
	if getProduct.Error != nil {
		return getProduct.Error
	}

	product.Stock -= cart.Quantity
	updateStock := config.DB.Where("id = ?", product.ID).Save(&product)
	if updateStock.Error != nil {
		return updateStock.Error
	}

	tx := config.DB.Create(&newOrderDetail)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func GetOrders(UserId int) (interface{}, error) {
	var orders []models.GetOrder

	check := config.DB.Where("orders.user_id = ?", UserId).Find(&models.Order{})
	if check.Error != nil {
		return nil, check.Error
	}

	tx := config.DB.Model(&models.Order{}).Select("orders.id, orders.created_at, orders.quantity, orders.total_price, orders.status, users.username AS buyer").Joins("inner join users on orders.user_id = users.id").Scan(&orders)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return nil, tx.Error
	}

	return orders, nil
}

func GetOrderDetails(orderId int) (interface{}, error) {
	var orderDetails []models.GetOrderDetails

	check := config.DB.Where("order_details.order_id = ?", orderId).Find(&models.OrderDetail{})
	if check.Error != nil {
		return nil, check.Error
	}

	tx := config.DB.Model(&models.OrderDetail{}).Select("order_details.id, order_details.created_at, orders.quantity, order_details.total_price, order_details.order_id").Joins("INNER JOIN orders on order_details.order_id = order.id").Scan(&orderDetails)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return nil, tx.Error
	}

	return orderDetails, nil
}
