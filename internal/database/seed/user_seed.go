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
		Username: "superadmin",
		Password: hashedPassword,
		FullName: "Super Admin",
		Email: "admin@zenithpay.com",
		Role: "admin",
		IsActive: true,
	}

	err = db.Where(model.User{Username: "superadmin"}).FirstOrCreate(&admin).Error
	if err != nil {
		log.Printf("Failed to seed admin: %v", err)
	} else {
		fmt.Println("✅ seeding admin successfully: Username: 'superadmin', password: 'savxpms141221'")
	}
}
