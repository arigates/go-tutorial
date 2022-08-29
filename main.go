package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Can't load configuration!")
	}

	// load db configuration
	DBName := os.Getenv("DB_NAME")
	DBUsername := os.Getenv("DB_USERNAME")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Makassar", DBUsername, DBPassword, DBName, DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// debug query
		//Logger: logger.Default.LogMode(logger.Info),
	})
}
