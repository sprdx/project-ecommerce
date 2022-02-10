package databases

import (
	"fmt"
	"project-ecommerce/config"
	"project-ecommerce/models"
)

func CreateUser(newUser *models.User) (interface{}, error) {
	// insert value of newUser into database
	tx := config.DB.Create(&newUser)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return newUser, nil
}

func LoginUser(user *models.User) (interface{}, error) {
	// Update value of user token into database
	tx := config.DB.Save(&user)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return nil, tx.Error
	}
	var data = struct {
		ID    uint
		Token string
	}{
		ID:    user.ID,
		Token: user.Token,
	}
	return data, nil
}

func GetUserById(userId int) (interface{}, error) {
	var user models.User
	tx := config.DB.Where("id = ?", userId).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	birthdate := fmt.Sprint(user.Birthdate)
	data := models.GetUser{
		ID:        user.ID,
		Name:      user.Username,
		Email:     user.Email,
		Birthdate: birthdate[:10],
		Gender:    user.Gender,
		Phone:     user.Phone,
		Photo:     user.Photo,
	}

	return data, nil
}

func GetUserByEmail(email string) (*models.User, int64) {
	var user models.User
	tx := config.DB.Where("email = ?", email).First(&user)
	if tx.RowsAffected == 1 {
		return &user, 1
	}
	return &user, 0
}

func GetUser(id int) (*models.User, int64) {
	var user models.User
	tx := config.DB.Where("id = ?", id).First(&user)
	if tx.RowsAffected == 1 {
		return &user, 1
	}
	return &user, 0
}

func UpdateUser(user *models.User) error {
	tx := config.DB.Save(&user)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return tx.Error
	}
	return nil
}
