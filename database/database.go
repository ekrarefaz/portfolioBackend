package database

import (
	"fmt"
	"github.com/ekrarefaz/portfolioBackend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))

	// make database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Perform Auto Migrate
	err = db.AutoMigrate(&models.Testimonial{}, &models.About{}, &models.Certification{}, &models.Project{})
	if err != nil {
		log.Fatalf("failed to migrate database %v", err)
	}

	DB = db
}
