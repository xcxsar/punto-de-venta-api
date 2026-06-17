package repository

import (
	"pos-api/internal/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (r *ProductRepository) Create(p *models.Product) error {
	return r.DB.Create(p).Error
}

func (r *ProductRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	err := r.DB.Find(&products).Error
	return products, err
}

func (r *ProductRepository) GetByID(id string) (models.Product, error) {
	var p models.Product
	err := r.DB.First(&p, id).Error
	return p, err
}

func (r *ProductRepository) Update(id string, p *models.Product) error {
	return r.DB.Model(&models.Product{}).Where("id = ?", id).Updates(p).Error
}

func (r *ProductRepository) Delete(id string) error {
	return r.DB.Delete(&models.Product{}, id).Error
}
