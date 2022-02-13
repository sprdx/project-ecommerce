package databases

import (
	"fmt"
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

	check := config.DB.Where("deleted_at IS NULL").Find(&models.Product{})
	if check.Error != nil {
		return nil, check.Error
	}

	tx := config.DB.Model(&models.Product{}).Select("products.id, products.product_name, products.category, products.price, products.stock, products.detail, products.rating, products.photo, users.username AS seller").Joins("inner join users on products.user_id = users.id").Order("products.id asc").Scan(&products)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return nil, tx.Error
	}

	return products, nil
}

func GetProductById(id int) (interface{}, error) {
	var product models.GetProduct
	check := config.DB.Where("deleted_at IS NULL").First(&models.Product{}, id)
	if check.Error != nil {
		return nil, check.Error
	}

	tx := config.DB.Model(&models.Product{}).Select("products.id, products.product_name, products.category, products.price, products.stock, products.detail, products.rating, products.photo, users.username AS seller").Joins("inner join users on products.user_id = users.id").Where("products.id = ?", id).Scan(&product)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return product, nil
}

func GetTheProduct(id int) (models.Product, error) {
	var product models.Product
	tx := config.DB.Where("deleted_at IS NULL").First(&product, id)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return product, tx.Error
	}
	fmt.Println(product)
	return product, nil
}

func UpdateProduct(product *models.Product) error {
	tx := config.DB.Save(&product)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return tx.Error
	}
	return nil
}

func DeleteProduct(product *models.Product) error {
	tx := config.DB.Delete(&product, product.ID)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return tx.Error
	}
	return nil
}
