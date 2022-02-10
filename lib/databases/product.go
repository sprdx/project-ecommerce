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
	var products []models.GetProduct
	tx := config.DB.Model(&models.Product{}).Select("products.id, products.product_name, products.category, products.price, products.stock, products.detail, products.rating, products.photo, users.username AS seller").Joins("inner join users on products.user_id = users.id").Scan(&products)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return products, nil
}

func GetProductById(id int) (interface{}, error) {
	var product models.GetProduct
	tx := config.DB.Model(&models.Product{}).Select("products.id, products.product_name, products.category, products.price, products.stock, products.detail, products.rating, products.photo, users.username AS seller").Joins("inner join users on products.user_id = users.id").Where("products.id = ? AND products.deleted_at IS NULL", id).Scan(&product)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return product, nil
}

func GetTheProduct(id int) (*models.Product, error) {
	var product models.Product
	tx := config.DB.First(&product, id)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return &product, tx.Error
	}

	return &product, nil
}

func UpdateProduct(product *models.Product) error {
	tx := config.DB.Save(&product)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return tx.Error
	}
	return nil
}
