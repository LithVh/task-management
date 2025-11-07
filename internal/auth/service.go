package auth

import (
	"fmt"
	"time"

	"task-management/internal/config"
	"task-management/internal/user"
	"task-management/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(dto *RegisterRequest) (*AuthResponse, error)
	Login(dto *LoginRequest) (*AuthResponse, error)
}

type authService struct {
	config *config.Config
	repo   *user.Repository
}

func NewAuthService(config *config.Config, repo *user.Repository) AuthService {
	return authService{config: config, repo: repo}
}

func (s authService) Register(dto *RegisterRequest) (*AuthResponse, error) {
	exists, err := s.repo.EmailAvailale(dto.Email)
	if err != nil {
		return nil, fmt.Errorf("Register: %v", err)
	}
	if exists {
		return nil, fmt.Errorf("email already registered")
	}

	pwHash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to generate password hash - Register: %v", err)
	}

	new := &user.User{
		Name:         dto.Name,
		PasswordHash: string(pwHash),
		Email:        dto.Email,
		CreatedAt:    time.Now(),
	}

	err = s.repo.Create(new)
	if err != nil {
		return nil, fmt.Errorf("failed to add user to DB - Register: %v", err)
	}

	tokenString, err := utils.CreateToken(s.config)
	if err != nil {
		return nil, fmt.Errorf("failed to create new token - Register: %v", err)
	}

	return &AuthResponse{
		Token: tokenString,
		User:  user.UserResponse{new.ID, new.Name, new.Email},
	}, nil

}

func (s authService) Login(dto *LoginRequest) (*AuthResponse, error) {

	userInfo, err := s.repo.FindByEmail(dto.Email)
	if err != nil {
		return nil, fmt.Errorf("Login: %v", err)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(userInfo.PasswordHash), []byte(dto.Password)); err != nil {
		return nil, fmt.Errorf("password doesnt match - Login: %v", err)
	}

	token, err := utils.CreateToken(s.config)
	if err != nil {
		return nil, fmt.Errorf("failed to create token - Login: %v", err)
	}

	return &AuthResponse{
		Token: token,
		User:  user.UserResponse{userInfo.ID, userInfo.Name, userInfo.Email},
	}, nil

}
