package model

import (
	"time"

	"github.com/google/uuid"
)

type TransactionItems struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	TransactionID uuid.UUID   `json:"transaction_id" gorm:"type:uuid;not null"`
	Transaction   Transaction `json:"transaction" gorm:"foreignKey:TransactionID;references:ID"`

	ProductID uuid.UUID `json:"product_id" gorm:"type:uuid;not null"`
	Product   Product   `json:"product" gorm:"foreignKey:ProductID;references:ID"`

	ProductName  string `json:"product_name" gorm:"not null"`
	ProductPrice int64  `json:"product_price" gorm:"not null"`

	Quantity int   `json:"quantity" gorm:"not null"`
	Subtotal int64 `json:"subtotal" gorm:"not null"` 

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
