package seed

import (
	"fmt"
	"log"

	"github.com/savanyv/zenith-pay/internal/model"
	"github.com/savanyv/zenith-pay/internal/utils/helpers"
	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB, bc helpers.BcryptHelper) {
	hashedPassword, err := bc.HashPassword("savxpms141221")
	if err != nil {
		log.Fatalf("Failed to hash password for admin: %v", err)
	}

	admin := model.User{
		Username: "savanyv",
		Password: hashedPassword,
		FullName: "Mochamad Saddam",
		Email: "savanyvv@zenithpay.com",
		Role: "admin",
		IsActive: true,
	}

	err = db.Where(model.User{Username: "savanyv"}).FirstOrCreate(&admin).Error
	if err != nil {
		log.Printf("Failed to seed admin: %v", err)
	} else {
		fmt.Println("✅ seeding admin successfully: Username: 'savanyv', password: 'savxpms141221'")
	}
}
