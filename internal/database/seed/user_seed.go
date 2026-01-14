package seed

import (
	"log"
	"os"

	"github.com/savanyv/zenith-pay/internal/model"
	"github.com/savanyv/zenith-pay/internal/utils/helpers"
	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB, bc helpers.BcryptHelper) {
	username := os.Getenv("ADMIN_USERNAME")
	password := os.Getenv("ADMIN_PASSWORD")
	email := os.Getenv("ADMIN_EMAIL")
	fullName := os.Getenv("ADMIN_FULL_NAME")

	if username == "" || password == "" {
		log.Println("ADMIN seed skipped (env not set)")
		return
	}

	hashedPassword, err := bc.HashPassword(password)
	if err != nil {
		log.Printf("failed to hash password: %v", err)
		return
	}

	var admin model.User
	err = db.Where("username = ?", username).First(&admin).Error
	if err == nil {
		log.Println("ADMIN already exists, skipping seed")
		return
	}

	if err != gorm.ErrRecordNotFound {
		log.Printf("Failed to check admin existence: %v", err)
		return
	}

	admin = model.User{
		Username: username,
		Password: hashedPassword,
		FullName: fullName,
		Email: email,
		Role: model.AdminRole,
		IsActive: true,
	}

	if err := db.Create(&admin).Error; err != nil {
		log.Printf("failed to seed admin: %v", err)
		return
	}

	log.Println("ADMIN seeded successfully")
}
