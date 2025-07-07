package services

import (
	"Shopping-cart/models"
	"Shopping-cart/repositories"
	"Shopping-cart/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	SignUp(input *models.User) error
	Login(email, password string) (string, error)
}

type userService struct {
	repo repositories.UserRepo
}

func NewUserService(repo repositories.UserRepo) UserService {
	return &userService{repo}
}

func (s *userService) SignUp(input *models.User) error {
	_, err := s.repo.CheckEmail(input.Email)
	if err == nil {
		return errors.New("email da ton tai")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	hashpw, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	input.Password = string(hashpw)
	return s.repo.SignUp(input)
}

func (s *userService) Login(email, password string) (string, error) {
	user, err := s.repo.CheckEmail(email)
	if err != nil {
		return "", errors.New("email khong ton tai")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("mat khau khong dung vui long nhap lai")
	}

	token, err := utils.GenerateJWT(user.Id)
	if err != nil {
		return "", nil
	}
	return token, nil
}
