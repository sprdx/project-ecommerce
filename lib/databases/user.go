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
	var user models.GetUser
	tx := config.DB.Select("id, username AS name, email, birthdate, gender, phone, photo").First(&models.User{}, userId).Scan(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	user.Birthdate = fmt.Sprint(user.Birthdate)
	user.Birthdate = user.Birthdate[:10]
	if user.Birthdate == "0001-01-01" {
		user.Birthdate = ""
	}

	return user, nil
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
