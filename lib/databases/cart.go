package databases

import (
	"fmt"
	"project-ecommerce/config"
	"project-ecommerce/models"
)

func CreateCart(newCart *models.Cart) (interface{}, error) {
	var cart *models.Cart
	check := config.DB.Where(&models.Cart{UserID: newCart.UserID, ProductID: newCart.ProductID}).First(&cart)
	if check.Error != nil {
		tx := config.DB.Create(&newCart)
		if tx.Error != nil {
			return nil, tx.Error
		}
		return newCart, nil
	}
	newCart.Quantity += cart.Quantity
	tx := config.DB.Where("id = ?", cart.ID).Save(&newCart)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return newCart, nil
}

func GetCart(userId int) (interface{}, error) {
	var carts []models.GetCart
	var cart []models.Cart
	fmt.Println("1")
	check := config.DB.Find(&cart)
	if check.RowsAffected == 0 {
		fmt.Println("a")
		return nil, check.Error
	}
	fmt.Println("2", cart)
	tx := config.DB.Model(&models.Cart{}).Select("carts.id, products.product_name, carts.quantity, (products.price * carts.quantity) AS total_price").Joins("INNER JOIN products on products.id = carts.product_id").Scan(&carts)
	if tx.Error != nil {
		fmt.Println("b")
		return nil, tx.Error
	}
	fmt.Println("3")
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

func GetDetailCart(id int) (uint, float64, error) {
	var cart models.Cart
	tx := config.DB.Where("id = ? AND deleted_at IS NULL", id).First(&cart)
	if tx.Error != nil {
		return 0, 0, tx.Error
	}
	return cart.Quantity, 0, nil
}
