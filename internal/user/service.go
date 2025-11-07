package user

import (
	"errors"

	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// GetProfile retrieves a user's profile by ID
func (s *Service) GetProfile(userID uuid.UUID) (*UserResponse, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return nil, err
	}
	return ToUserResponse(user), nil
}

// UpdateProfile updates a user's profile
func (s *Service) UpdateProfile(userID uuid.UUID, dto *UpdateUserDTO) (*UserResponse, error) {
	// Get existing user
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if dto.Name != "" {
		user.Name = dto.Name
	}

	if dto.Email != "" {
		// Check if email is already in use by another user
		exists, err := s.repo.EmailExistsExcludingUser(dto.Email, userID)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, errors.New("email already in use")
		}
		user.Email = dto.Email
	}

	// Save updates
	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return ToUserResponse(user), nil
}
