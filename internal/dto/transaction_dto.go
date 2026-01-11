package dtos

import "time"

type TransactionItemRequest struct {
	ProductID string `json:"product_id" validate:"required,uuid"`
	Quantity  int    `json:"quantity" validate:"required,gt=0"`
}

type TransactionRequest struct {
	PaymentMethod string                   `json:"payment_method" validate:"required"`
	PaymentAmount float64                  `json:"payment_amount" validate:"required,gt=0"`
	Items         []TransactionItemRequest `json:"items" validate:"required,dive"`
}

type TransactionItemResponse struct {
	ProductID    string  `json:"product_id"`
	ProductName  string  `json:"product_name"`
	ProductPrice float64 `json:"product_price"`
	Quantity     int     `json:"quantity"`
	SubTotal     float64 `json:"sub_total"`
}

type TransactionResponse struct {
	ID              string                    `json:"id"`
	UserID          string                    `json:"user_id"`
	TransactionDate time.Time                 `json:"transaction_date"`
	PaymentMethod   string                    `json:"payment_method"`
	TotalAmount     float64                   `json:"total_amount"`
	PaymentAmount   float64                   `json:"payment_amount"`
	ChangeAmount    float64                   `json:"change_amount"`
	Items           []TransactionItemResponse `json:"items,omitempty"`
}
