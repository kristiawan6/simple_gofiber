package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	
)

var DB *gorm.DB

func InitDB() {
	url := "DATABASE_URL"
	// url := os.Getenv("URL") 
	var err error
	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

}
