package repositories

import (
	"Shopping-cart/models"
	"errors"

	"gorm.io/gorm"
)

type CartRepo interface {
	GetCartItem(userID, productID uint) (*models.CartItem, error)
	CreateCartItem(item *models.CartItem) error
	UpdateCartItem(item *models.CartItem) error
	GetProduct(productID uint, product *models.Product) error
	ListItems(userID uint, page int, limit int) (*models.ListCart, error)
	DeleteItem(productID uint) error
	DeleteMany(IdsProduct []uint) error
}

type cartRepo struct {
	db *gorm.DB
}

func NewCartRepo(db *gorm.DB) CartRepo {
	return &cartRepo{db}
}

func (r *cartRepo) GetCartItem(userID, productID uint) (*models.CartItem, error) {
	var item models.CartItem
	err := r.db.Where("user_id = ? AND product_id = ?", userID, productID).First(&item).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &item, nil

}

func (r *cartRepo) CreateCartItem(item *models.CartItem) error {
	return r.db.Create(&item).Debug().Error
}

func (r *cartRepo) UpdateCartItem(item *models.CartItem) error {
	return r.db.Save(&item).Error
}

func (r *cartRepo) GetProduct(productID uint, product *models.Product) error {
	return r.db.Where("id = ?", productID).First(&product).Error
}

func (r *cartRepo) ListItems(userID uint, page int, limit int) (*models.ListCart, error) {
	var items []models.CartItem
	offset := (page - 1) * limit
	if err := r.db.Where("user_id = ?", userID).Order("id desc").Offset(offset).Limit(limit).Find(&items).Error; err != nil {
		return nil, err
	}
	var total float64
	for _, v := range items {
		total += v.Price
	}
	listCart := &models.ListCart{
		IDCart:     userID,
		Items:      items,
		TotalPrice: total,
	}
	return listCart, nil
}

func (r *cartRepo) DeleteItem(productID uint) error {
	return r.db.Table("cart_items").Where(" product_id = ?", productID).Delete(nil).Error
}

func (r *cartRepo) DeleteMany(IdsProduct []uint) error {
	return r.db.Table("cart_items").Where("product_id IN ?", IdsProduct).Delete(nil).Error
}
