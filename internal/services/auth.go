package services

import (
	"backend/internal/models"
	"backend/internal/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{UserRepo: userRepo}
}

func (s *AuthService) Register(username, password string) (int64, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return s.UserRepo.CreateUser(username, string(hashedPassword))
}

func (s *AuthService) Login(username, password string) (*models.User, error) {
	user, err := s.UserRepo.GetUserByUsername(username)
	if err != nil {
		return nil, errors.New("user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}
	return user, nil
}
