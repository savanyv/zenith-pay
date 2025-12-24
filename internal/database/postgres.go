package database

import (
	"fmt"
	"log"

	"github.com/savanyv/zenith-pay/config"
	"github.com/savanyv/zenith-pay/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to database:", err)
			return nil, err
		}

		if err := db.AutoMigrate(
			&model.User{},
			&model.Category{},
			&model.Transaction{},
			&model.TransactionItems{},
			&model.Product{},
		); err != nil {
			log.Fatal("Failed to migrate database:", err)
			return nil, err
		}

		DB = db

		log.Println("Database connection successfully")
		return db, nil
}
