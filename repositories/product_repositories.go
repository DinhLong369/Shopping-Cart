package repositories

import (
	"Shopping-cart/models"

	"gorm.io/gorm"
)

type ProductRepo interface {
	CreateProduct(input *models.Product) error
	GetByID(id int) (*models.Product, error)
	UpdateProduct(id int, input *models.CreateProduct) error
	DeleteProduct(id int) error
	ListProduct(page int, limit int) ([]models.Product, int64, error)
	DeleteMany(ids []uint) error
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	return &productRepo{db}
}

func (r *productRepo) CreateProduct(input *models.Product) error {
	return r.db.Create(input).Error
}
func (r *productRepo) GetByID(id int) (*models.Product, error) {
	var result models.Product
	err := r.db.Table("products").Where("id = ?", id).First(&result).Error
	return &result, err
}

func (r *productRepo) UpdateProduct(id int, input *models.CreateProduct) error {
	return r.db.Table("products").Where("id = ?", id).Updates(&input).Error
}

func (r *productRepo) DeleteProduct(id int) error {
	var p models.Product
	return r.db.Where("id = ?", id).Delete(&p).Debug().Error
}
func (r *productRepo) ListProduct(page int, limit int) ([]models.Product, int64, error) {
	var items []models.Product
	var total int64
	offset := (page - 1) * limit
	r.db.Table("products").Count(&total)
	err := r.db.Order("id desc").Offset(offset).Limit(limit).Find(&items).Error
	return items, total, err
}

func (r *productRepo) DeleteMany(ids []uint) error {
	var p models.Product
	return r.db.Table("products").Where("id IN ?", ids).Delete(&p).Error
}
