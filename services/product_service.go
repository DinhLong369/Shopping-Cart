package services

import (
	"Shopping-cart/models"
	"Shopping-cart/repositories"
)

type ProductService interface {
	CreateProduct(input *models.Product) error
	GetByID(id int) (*models.Product, error)
	UpdateProduct(id int, input *models.CreateProduct) error
	DeleteProduct(id int) error
	ListProduct(page int, limit int) ([]models.Product, int64, error)
	DeleteMany(ids []uint) error
}
type productService struct {
	repo repositories.ProductRepo
}

func NewProductService(repo repositories.ProductRepo) ProductService {
	return &productService{repo}
}

func (s *productService) CreateProduct(input *models.Product) error {
	return s.repo.CreateProduct(input)
}
func (s *productService) GetByID(id int) (*models.Product, error) {
	return s.repo.GetByID(id)
}
func (s *productService) UpdateProduct(id int, input *models.CreateProduct) error {
	return s.repo.UpdateProduct(id, input)
}
func (s *productService) DeleteProduct(id int) error {
	return s.repo.DeleteProduct(id)
}

func (s *productService) ListProduct(page int, limit int) ([]models.Product, int64, error) {
	return s.repo.ListProduct(page, limit)
}

func (s *productService) DeleteMany(ids []uint) error {
	return s.repo.DeleteMany(ids)
}
