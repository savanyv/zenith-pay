package repository

import (
	"github.com/savanyv/zenith-pay/internal/model"
	"gorm.io/gorm"
)

type TransactionItemRepository interface {
	CreateMany(tx *gorm.DB, items []model.TransactionItems) error
	FindByTransactionID(transactionID string) ([]model.TransactionItems, error)
}

type transactionItemRepository struct {
	db *gorm.DB
}

func NewTransactionItemRepository(db *gorm.DB) TransactionItemRepository {
	return &transactionItemRepository{db: db}
}

// CreateMany digunakan untuk menyimpan banyak item sekaligus dalam satu transaksi.
func (r *transactionItemRepository) CreateMany(tx *gorm.DB, items []model.TransactionItems) error {
	db := r.db
	if tx != nil {
		db = tx
	}
	return db.Create(&items).Error
}

func (r *transactionItemRepository) FindByTransactionID(transactionID string) ([]model.TransactionItems, error) {
	var items []model.TransactionItems
	if err := r.db.Where("transaction_id = ?", transactionID).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
