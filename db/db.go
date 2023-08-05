package db

import (
	"everything-is-twenty/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Init(dbUrl string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
