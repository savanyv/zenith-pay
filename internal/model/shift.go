package model

import (
	"time"

	"github.com/google/uuid"
)

type ShiftStatus string

const (
	ShiftOpen ShiftStatus = "open"
	ShiftClose ShiftStatus = "closed"
)

func (s ShiftStatus) IsValid() bool {
	switch s {
	case ShiftOpen, ShiftClose:
		return true
	default:
		return false
	}
}

type Shift struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CashierID uuid.UUID `json:"cashier_id" gorm:"type:uuid;not null"`
	Status ShiftStatus `json:"status" gorm:"not null"`
	OpeningBalance int64 `json:"opening_balance" gorm:"not null; default: 0"`
	ClosingBalance *int64 `json:"closing_balance" gorm:"not null; default: 0"`
	OpenedAt time.Time `json:"opened_at" gorm:"not null"`
	ClosedAt time.Time `json:"closed_at" gorm:"not null"`
}
