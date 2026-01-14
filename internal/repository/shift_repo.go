package repository

import (
	"errors"

	"github.com/savanyv/zenith-pay/internal/model"
	"gorm.io/gorm"
)

type ShiftRepository interface {
	Create(shift *model.Shift) error
	FindActiveShiftByCashier(cashierID string) (*model.Shift, error)
	FindByID(ID string) (*model.Shift, error)
	CloseShift(shift *model.Shift) error
}

type shiftRepository struct {
	db *gorm.DB
}

func NewShiftRepository(db *gorm.DB) ShiftRepository {
	return &shiftRepository{
		db: db,
	}
}

func (r *shiftRepository) Create(shift *model.Shift) error {
	if err := r.db.Create(&shift).Error; err != nil {
		return err
	}

	return nil
}

func (r *shiftRepository) FindActiveShiftByCashier(cashierID string) (*model.Shift, error) {
	var shift model.Shift
	err := r.db.Where("cashier_id = ? AND status = ?", cashierID, model.ShiftOpen).First(&shift).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &shift, nil
}

func (r *shiftRepository) FindByID(ID string) (*model.Shift, error) {
	var shift model.Shift
	err := r.db.First(&shift, "id = ?", ID). Error
	if err != nil {
		return nil, err
	}

	return &shift, nil
}

func (r *shiftRepository) CloseShift(shift *model.Shift) error {
	return r.db.Save(shift).Error
}
