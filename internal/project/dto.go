package project

import (
	"time"

	"github.com/google/uuid"
)

type CreateProjectRequest struct {
	Name        string `json:"name" binding:"required,min=1,max=255"`
	Description string `json:"description" binding:"omitempty,max=255"`
}

type UpdateProjectRequest struct {
	Name        string `json:"name" binding:"omitempty,min=1,max=255"`
	Description string `json:"description" binding:"omitempty,max=255"`
}

type ProjectResponse struct {
	ID          int64     `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ToProjectResponse(project *Project) *ProjectResponse {
	return &ProjectResponse{
		ID:          project.ID,
		UserID:      project.OwnerID,
		Name:        project.Name,
		Description: project.Description,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
	}
}

func ToProjectResponseList(projects []*Project) []*ProjectResponse {
	responses := make([]*ProjectResponse, len(projects))
	for i, project := range projects {
		responses[i] = ToProjectResponse(project)
	}
	return responses
}
