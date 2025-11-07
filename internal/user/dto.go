package user

import "github.com/google/uuid"

// UpdateUserDTO represents the request body for updating user profile
type UpdateUserDTO struct {
	Name  string `json:"name" binding:"omitempty,min=2,max=255"`
	Email string `json:"email" binding:"omitempty,email,max=255"`
}

type UserResponse struct {
	ID uuid.UUID `json:"id"`
	Name string	`json:"name"`
	Email string	`json:"email"`
}

// ToUserResponse converts a User model to UserResponse
func ToUserResponse(user *User) *UserResponse {
	return &UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
