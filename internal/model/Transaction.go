package model

import (
	"time"

	"github.com/google/uuid"
)

type PaymentMethod string

const (
	Cash PaymentMethod = "Cash"
	Debit PaymentMethod = "Debit"
	Qris PaymentMethod = "Qris"
)

func (p PaymentMethod) IsValid() bool {
	switch p {
	case Cash, Debit, Qris:
		return true
	default:
		return false
	}
}

type Transaction struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	User User `json:"user" gorm:"foreignKey:UserID;references:ID"`
	TransactionDate time.Time `json:"transaction_date" gorm:"not null"`
	PaymentMethod PaymentMethod `json:"payment_method" gorm:"not null"`
	TotalAmount float64 `json:"total_amount" gorm:"not null"`
	PaymentAmount float64 `json:"payment_amount" gorm:"not null"`
	ChangeAmount float64 `json:"change_amount" gorm:"not null"`
	TransactionItems []TransactionItems `json:"transaction_items" gorm:"foreignKey:TransactionID;references:ID"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
