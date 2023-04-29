package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func CreateDB() *gorm.DB {
	dsn := AppConfig.PostgresConfig.GetPostgresConnStr()
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Database")
	}
	return database
}
