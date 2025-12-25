package repository

import (
	"github.com/savanyv/zenith-pay/internal/model"
	"gorm.io/gorm"
)

type UserRespository interface {
	Create(user *model.User) error
	GetByUsername(username string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRespository {
	return &userRepository{
		db: db,
	}
}

// Create inserts a new user into the database
func (r *userRepository) Create(user *model.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

// GetByUsername retrieves a user by their username
func (r *userRepository) GetByUsername(username string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
