package database

import (
	"log"
	"os"
	"telephone/internal/modules"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectToDatabase() {
	// create a connection for database
	dns := os.Getenv("DNS")
	database, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatal("error cont connect to database", err)
	}

	Database = database
	// auto migrate
	Database.AutoMigrate(&modules.User{})
}
