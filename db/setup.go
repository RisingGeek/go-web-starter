package db

import (
	"fmt"

	"github.com/risinggeek/golang-starter/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Setup export
func Setup() *gorm.DB {
	dsn := "root:mysql@tcp(localhost:3306)/todo_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error connecting to database")
	} else {
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Todo{})

		fmt.Println("Successfully connected to database")
	}
	return db
}
