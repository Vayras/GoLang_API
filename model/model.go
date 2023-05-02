package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Person struct {
	Id       int    `json:"id" gorm:"primary_key,auto_increment"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(mysql.Open("root:roottoor@/testdb"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}
	DB.AutoMigrate(&Person{})
}
