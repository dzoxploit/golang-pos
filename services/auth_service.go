package services

import (
	"go-pos/models"
	"go-pos/repositories"
	"go-pos/utils"

	"github.com/jinzhu/gorm"
)

type AuthService struct {
	userRepository *repositories.UserRepository
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		userRepository: repositories.NewUserRepository(db),
	}
}

func (s *AuthService) Register(user *models.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	return s.userRepository.CreateUser(user)
}

func (s *AuthService) Login(username, password string) (string, error) {
	user, err := s.userRepository.GetUserByUsername(username)
	if err != nil {
		return "", err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", utils.ErrInvalidCredentials
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
