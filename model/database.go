package model

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	dsn := "host=localhost user=root password=password dbname=postgres port=5555 sslmode=disable TimeZone=Asia/Tokyo"
	var err error
	if Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		panic("failed to connect database")
	}
	log.Println("database connected!")
	Db.AutoMigrate(&BlogEntity{})
	log.Println("migrated")
}
