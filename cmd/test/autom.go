package main

import (
	"fmt"

	"github.com/rn-consider/gormgenex/dao/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database:", err)
		return
	}

	// AutoMigrate the User model
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Println("failed to migrate database:", err)
	}
}
