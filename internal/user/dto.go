package user

import "github.com/google/uuid"

type UpdateUserRequest struct {
	Name string `json:"name" binding:"required,min=1,max=255"`
	// Email string `json:"email" binding:"omitempty,email,max=255"`
}

type UserResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

