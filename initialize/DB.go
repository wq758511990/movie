package initialize

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"todo-pro/globals"
	"todo-pro/models"
)

func InitDB() {
	db, err := gorm.Open("mysql", "user:123456@/com.movie?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Movie{})
	db.AutoMigrate(&models.Category{})

	globals.DB = db
}
