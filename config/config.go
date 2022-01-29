package config

import (
	"project-ecommerce/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Define GORM Database
var DB *gorm.DB

//Declare function to connect database
func InitDB() {
	//Set data source that will be used
	connection := "root:qwerty@tcp(127.0.0.1:3306)/alta_ecommerce?charset=utf8mb4&parseTime=True&loc=Local"

	//Initialize DB session
	var err error
	DB, err = gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	//Migrate the database schema
	InitMigration()
}

// Declare function to auto-migrate the schema
func InitMigration() {
	// DB.Migrator().DropTable(&models.User{}, &models.Product{}, &models.Cart{})
	DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Cart{})
}
