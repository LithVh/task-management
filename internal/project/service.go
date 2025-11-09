package project

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	List(userID uuid.UUID) ([]*ProjectResponse, error)
	Create(userID uuid.UUID, dto *CreateProjectRequest) (*ProjectResponse, error)
	GetByID(projectID int64, userID uuid.UUID) (*ProjectResponse, error)
	Update(projectID int64, userID uuid.UUID, dto *UpdateProjectRequest) (*ProjectResponse, error)
	Delete(projectID int64, userID uuid.UUID) error
}

type service struct {
	repo *Repository
}

func NewService(repo *Repository) Service {
	return &service{repo: repo}
}

func (s *service) List(userID uuid.UUID) ([]*ProjectResponse, error) {
	projects, err := s.repo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}
	return ToProjectResponseList(projects), nil
}

func (s *service) Create(userID uuid.UUID, dto *CreateProjectRequest) (*ProjectResponse, error) {
	now := time.Now()
	project := &Project{
		UserID:      userID,
		Name:        dto.Name,
		Description: dto.Description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := s.repo.Create(project); err != nil {
		return nil, err
	}

	return ToProjectResponse(project), nil
}

func (s *service) GetByID(projectID int64, userID uuid.UUID) (*ProjectResponse, error) {
	project, err := s.repo.FindByID(projectID)
	if err != nil {
		return nil, err
	}

	//only owner can access
	if project.UserID != userID {
		return nil, errors.New("unauthorized: you don't own this project")
	}

	return ToProjectResponse(project), nil
}

func (s *service) Update(projectID int64, userID uuid.UUID, dto *UpdateProjectRequest) (*ProjectResponse, error) {
	project, err := s.repo.FindByID(projectID)
	if err != nil {
		return nil, err
	}

	if project.UserID != userID {
		return nil, errors.New("unauthorized: you don't own this project")
	}

	if dto.Name != "" {
		project.Name = dto.Name
	}
	if dto.Description != "" {
		project.Description = dto.Description
	}
	project.UpdatedAt = time.Now()

	if err := s.repo.Update(project); err != nil {
		return nil, err
	}

	return ToProjectResponse(project), nil
}

func (s *service) Delete(projectID int64, userID uuid.UUID) error {
	project, err := s.repo.FindByID(projectID)
	if err != nil {
		return err
	}

	if project.UserID != userID {
		return errors.New("unauthorized: you don't own this project")
	}

	return s.repo.Delete(projectID)
}
