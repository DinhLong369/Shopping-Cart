package services

import (
	"Shopping-cart/models"
	"Shopping-cart/repositories"
)

type CartService interface {
	AddToCart(userID, productID, quantity uint) error
	ListItems(userID uint, page int, limit int) (*models.ListCart, error)
	DeleteItem(productID uint) error
	DeleteMany(IdsProduct []uint) error
	UpdateCartItem(productID, userID, quantity uint) error
}

type cartService struct {
	repo repositories.CartRepo
}

func NewCartService(repo repositories.CartRepo) CartService {
	return &cartService{repo}
}

func (s *cartService) AddToCart(userID, productID, quantity uint) error {
	var product models.Product
	err := s.repo.GetProduct(productID, &product)
	if err != nil {
		return err
	}
	existItem, err := s.repo.GetCartItem(userID, productID)
	if err != nil {
		return err
	}
	if existItem != nil {
		existItem.Quantity += quantity
		existItem.Price = product.Price * float64(existItem.Quantity)
		return s.repo.UpdateCartItem(existItem)
	}

	newItem := &models.CartItem{
		Quantity:  quantity,
		UserID:    userID,
		ProductID: productID,
		Price:     product.Price * float64(quantity),
	}
	return s.repo.CreateCartItem(newItem)
}

func (s *cartService) ListItems(userID uint, page int, limit int) (*models.ListCart, error) {
	return s.repo.ListItems(userID, page, limit)
}

func (s *cartService) DeleteItem(productID uint) error {
	return s.repo.DeleteItem(productID)
}

func (s *cartService) DeleteMany(IdsProduct []uint) error {
	return s.repo.DeleteMany(IdsProduct)
}

func (s *cartService) UpdateCartItem(productID, userID, quantity uint) error {
	var product models.Product
	err := s.repo.GetProduct(productID, &product)
	if err != nil {
		return err
	}
	existItem, err := s.repo.GetCartItem(userID, productID)
	if err != nil {
		return err
	}

	if existItem != nil {
		existItem.Quantity = quantity
		existItem.Price = product.Price * float64(existItem.Quantity)
	}
	return s.repo.UpdateCartItem(existItem)
}
