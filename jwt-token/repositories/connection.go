package repositories

import (
	"jwt-token/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Enter Password for your MySQL database
func Connect() {
	connection, err := gorm.Open(mysql.Open("root:passwordHere@/jwt-token"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
