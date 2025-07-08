package services

import (
	"Shopping-cart/models"
	"Shopping-cart/repositories"
	"Shopping-cart/utils"
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	SignUp(input *models.RequestSignUp) error
	Login(email, password string) (string, error)
}

type userService struct {
	repo repositories.UserRepo
}

func NewUserService(repo repositories.UserRepo) UserService {
	return &userService{repo}
}

func (s *userService) SignUp(input *models.RequestSignUp) error {
	_, err := s.repo.CheckEmail(input.Email)
	if err == nil {
		return errors.New("email da ton tai")
	}
	fmt.Println("day la", input.Password)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	hashpw, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	input.Password = string(hashpw)
	return s.repo.SignUp(input)
}

func (s *userService) Login(email, password string) (string, error) {
	user, err := s.repo.CheckEmail(email)
	if err != nil {
		return "", errors.New("email khong ton tai")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("mat khau khong dung vui long nhap lai")
	}

	token, err := utils.GenerateJWT(user.Id)
	if err != nil {
		return "", errors.New("khong tao duoc token")
	}
	return token, nil
}
