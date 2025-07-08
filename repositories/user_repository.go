package repositories

import (
	"Shopping-cart/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	CheckEmail(email string) (*models.User, error)
	SignUp(input *models.RequestSignUp) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db}
}

func (r *userRepo) CheckEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepo) SignUp(input *models.RequestSignUp) error {
	return r.db.Table("users").Create(input).Error
}
