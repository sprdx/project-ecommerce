package databases

import (
	"project-ecommerce/config"
	"project-ecommerce/models"
)

func CreateProduct(newProduct *models.Product) (interface{}, error) {
	tx := config.DB.Create(&newProduct)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return newProduct, nil
}

func GetAllProducts() (interface{}, error) {
	var products *[]models.Product
	tx := config.DB.Find(&products)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return products, nil
}

func GetProductById(id int) (interface{}, error) {
	var product models.Product
	tx := config.DB.Where("id = ?", id).First(&product)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return product, nil
}
