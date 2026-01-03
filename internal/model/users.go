package model

import (
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	AdminRole Role = "admin"
	CashierRole Role = "cashier"
)

func (r Role) IsValid() bool {
	switch r {
	case AdminRole, CashierRole:
		return true
	default:
		return false
	}
}

type User struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Username string   `json:"username" gorm:"unique;not null"`
	Password string   `json:"password" gorm:"not null"`
	FullName string   `json:"full_name" gorm:"not null"`
	Email    string   `json:"email" gorm:"unique;not null"`
	Role Role   `json:"role" gorm:"not null"`
	IsActive bool    `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
