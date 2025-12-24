package model

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	User User `json:"user" gorm:"foreignKey:UserID;references:ID"`
	TransactionDate time.Time `json:"transaction_date" gorm:"not null"`
	PaymentMethod string `json:"payment_method" gorm:"not null"`
	TotalAmount float64 `json:"total_amount" gorm:"not null"`
	ChangeAmount float64 `json:"change_amount" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
