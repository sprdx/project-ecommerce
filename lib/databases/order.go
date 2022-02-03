package databases

import (
	"project-ecommerce/config"
	"project-ecommerce/models"
)

func CreateOrder(newOrder *models.Order) error {
	tx := config.DB.Create(&newOrder)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
