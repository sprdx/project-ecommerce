package databases

import "project-ecommerce/models"

func CreateCart(id int, cart *models.Cart) (interface{}, error) {
	product, _ := GetProductById(id)
}
