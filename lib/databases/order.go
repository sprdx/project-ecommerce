package databases

import (
	"project-ecommerce/config"
	"project-ecommerce/models"
)

func CreateOrder(newPayment models.Payment) error {
	var order models.Order
	order.Payment = newPayment
	tx := config.DB.Create(&order)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
