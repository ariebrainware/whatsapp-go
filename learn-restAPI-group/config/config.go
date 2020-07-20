package config

// go get github.com/jinzhu/gorm.
import (
	"github.com/jinzhu/gorm"
)

// db init create connection to database.
func dbInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:Hallo123$@tcp(127.0.0.1:3306)/chat?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to Connect to Database!")
	}
	db.AutoMigrate(structs.User{})
	return db
}
