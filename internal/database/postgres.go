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
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if cfg.AppEnv == "development" {
		if err := db.AutoMigrate(
			&model.User{},
			&model.Category{},
			&model.Product{},
			&model.Transaction{},
			&model.TransactionItems{},
			&model.Shift{},
		); err != nil {
			log.Fatal(err)
		}
	}

	DB = db
	log.Println("Database connected")
	return db, nil
}
