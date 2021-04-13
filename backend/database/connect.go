package database

import (
	"../models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect() {


//	db, err := gorm.Open(mysql.Open(os.Getenv("USER")+":"+os.Getenv("PASSWORD")+"@/"+os.Getenv("DBNAME")), &gorm.Config{})
	db,err := gorm.Open(mysql.Open("root:Rollout3@/mars"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the DB")
	}

	DB = db

	db.AutoMigrate(&models.User{}, &models.Role{})
	log.Println(db)
}
