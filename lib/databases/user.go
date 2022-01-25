package databases

import (
	"project-ecommerce/config"
	"project-ecommerce/middlewares"
	"project-ecommerce/models"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(newUser *models.User) (interface{}, error) {
	// insert value of newUser into database
	tx := config.DB.Create(&newUser)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return newUser, nil
}

func LoginUser(userData *models.User) (interface{}, error) {
	var user models.User
	// Check if user's login email is exist in database
	tx := config.DB.Where("email = ?", userData.Email).First(&user)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return nil, tx.Error
	}

	// Check if inputed password is match to password in database
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData.Password))
	if err != nil {
		return nil, err
	}

	// Generate token by using user's ID
	user.Token, _ = middlewares.CreateToken(int(user.ID))

	// Update value of user token into database
	tx2 := config.DB.Save(&user)
	if tx2.Error != nil || tx.RowsAffected == 0 {
		return nil, tx.Error
	}

	return user.Token, nil
}

func GetUserById(userId int) (interface{}, error) {
	var user models.User
	tx := config.DB.Where("id = ?", userId).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	userData := models.GetUser{
		ID:        user.ID,
		Name:      user.User_name,
		Email:     user.Email,
		Birthdate: user.Birthdate,
		Gender:    user.Gender,
		Phone:     user.Phone,
		Photo:     user.Photo,
	}

	return userData, nil
}
