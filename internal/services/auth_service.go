package services

import (
	"time"

	"github.com/Hattaseakhiaw/sre-user-management/backend/internal/models"
	"github.com/Hattaseakhiaw/sre-user-management/backend/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(input *models.RegisterRequest) (*models.User, error)
	Login(input *models.LoginRequest) (string, error)
}

type authService struct {
	userRepo repository.UserRepository
	jwtKey   []byte
}

func NewAuthService(userRepo repository.UserRepository, jwtKey []byte) AuthService {
	return &authService{
		userRepo: userRepo,
		jwtKey:   jwtKey,
	}
}

func (s *authService) Register(input *models.RegisterRequest) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	err = s.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authService) Login(input *models.LoginRequest) (string, error) {
	user, err := s.userRepo.GetUserByEmail(input.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return "", err // password ไม่ถูกต้อง
	}

	// สร้าง JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // อายุ token 72 ชั่วโมง
	})

	tokenString, err := token.SignedString(s.jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
