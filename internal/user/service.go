package user

import (
	"fmt"

	"github.com/google/uuid"
)

type Service interface {
	GetProfile(id uuid.UUID) (*UserResponse, error)
	UpdateProfile(id uuid.UUID, dto *UpdateUserRequest) (*UserResponse, error)
}

type service struct {
	repo *Repository
}

func NewService(repo *Repository) Service {
	return &service{repo: repo}
}

func (service *service) GetProfile(id uuid.UUID) (*UserResponse, error) {
	user, err := service.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("GetProfile: %v", err)
	}

	return &UserResponse{user.ID, user.Name, user.Email}, nil
}

func (service *service) UpdateProfile(id uuid.UUID, dto *UpdateUserRequest) (*UserResponse, error) {
	user, err := service.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("GetProfile: %v", err)
	}

	user.Name = dto.Name
	err = service.repo.Update(user)
	if err != nil {
		return nil, fmt.Errorf("failed to update profile - UpdateProfile: %v", err)
	}

	return &UserResponse{user.ID, user.Name, user.Email}, nil
}

