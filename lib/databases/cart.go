package databases

import (
	"project-ecommerce/config"
	"project-ecommerce/models"
)

func CreateCart(newCart *models.Cart) (interface{}, error) {
	var product models.Product
	tx := config.DB.Where("id = ?", newCart.ProductID).First(&product)
	if tx.Error != nil {
		return nil, tx.Error
	}

	product = models.Product{
		Price: product.Price,
	}

	newCart.Total_Price = product.Price * float64(newCart.Quantity)

	tx2 := config.DB.Create(&newCart)
	if tx2.Error != nil {
		return nil, tx2.Error
	}

	return newCart, nil
}

func GetCart(id int) (interface{}, error) {
	var carts []models.GetCart
	tx := config.DB.Table("carts").Select("products.product_name, carts.quantity, carts.total_price").Joins("INNER JOIN products on products.id = carts.product_id").Where("carts.user_id = ? AND carts.deleted_at IS NULL", id).Scan(&carts)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return carts, nil
}

func DeleteCart(userId int, productId int) error {
	var cart models.Cart
	tx := config.DB.Where(&models.Cart{UserID: uint(userId), ProductID: uint(productId)}).Find(&cart)
	if tx.Error != nil {
		return tx.Error
	}

	tx2 := config.DB.Where("id = ?", cart.ID).Delete(&models.Cart{})
	if tx2.Error != nil {
		return tx2.Error
	}

	return nil
}
