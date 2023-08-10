package database

import (
	"fmt"

	"example.com/blog/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:123456789@tcp(127.0.0.1:3306)/go_blog_project?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	} else {
		fmt.Println("Connect success")
	}
	DB = db
	db.AutoMigrate(
		&models.User{},
		&models.Blog{},
	)
}