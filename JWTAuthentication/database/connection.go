package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	_, err := gorm.Open(mysql.Open("root:603392@/JWTAuthentication"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}
}
