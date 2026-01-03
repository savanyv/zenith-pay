package repository

import (
	"github.com/savanyv/zenith-pay/internal/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *model.Product) error
	FindByID(id string) (*model.Product, error)
	FindBySKU(sku string) (*model.Product, error)
	FindByName(name string) (*model.Product, error)
	FindAll() ([]*model.Product, error)
	Update(product *model.Product) error
	Delete(id string) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) Create(product *model.Product) error {
	if err := r.db.Create(product).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepository) FindByID(id string) (*model.Product, error) {
	var product model.Product
	if err := r.db.Preload("Category").Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) FindBySKU(sku string) (*model.Product, error) {
	var product model.Product
	if err := r.db.Preload("Category").Where("sku = ?", sku).First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) FindByName(name string) (*model.Product, error) {
	var product model.Product
	if err := r.db.Preload("Category").Where("name = ?", name).First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) FindAll() ([]*model.Product, error) {
	var products []*model.Product
	if err := r.db.Preload("Category").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepository) Update(product *model.Product) error {
	if err := r.db.Save(product).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepository) Delete(id string) error {
	if err := r.db.Delete(&model.Product{}, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}
