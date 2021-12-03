package app

import (
	"fmt"
	"github.com/nunonano/hacktiv-assignment2/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitDB() *gorm.DB{
	dsn := "root:secret@tcp(127.0.0.1:3306)/orders_by?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		log.Panic("Database error", err.Error())
	}

	log.Println("Database connected")
	db.Debug().AutoMigrate(model.Orders{}, model.Items{})

	return db
}