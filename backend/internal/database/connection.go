package database

import (
	"github.com/AIFuzi/hakuro-shop/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func New() {
	conn, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")))
	if err != nil {
		log.Fatal("Fatal error: Failed to connect database")
	}

	DB = conn

	err = conn.AutoMigrate(models.User{}, models.Types{}, models.Product{})
	if err != nil {
		log.Fatal("Fatal error: Failed to migrate models")
	}
}
