package project

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	List(userID uuid.UUID) ([]*ProjectResponse, error)
	Create(userID uuid.UUID, dto *CreateProjectRequest) (*ProjectResponse, error)
	GetByID(projectID int64, userID uuid.UUID) (*ProjectResponse, error)
	Update(projectID int64, userID uuid.UUID, dto *UpdateProjectRequest) (*ProjectResponse, error)
	Delete(projectID int64, userID uuid.UUID) error
	AddUser(projectID int64, userID uuid.UUID, targetID uuid.UUID) error
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
		OwnerID:     userID,
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
	if project.OwnerID != userID {
		return nil, fmt.Errorf("unauthorized: you don't own this project")
	}

	return ToProjectResponse(project), nil
}

func (s *service) Update(projectID int64, userID uuid.UUID, dto *UpdateProjectRequest) (*ProjectResponse, error) {
	project, err := s.repo.FindByID(projectID)
	if err != nil {
		return nil, err
	}

	if project.OwnerID != userID {
		return nil, fmt.Errorf("unauthorized: you don't own this project")
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

	if project.OwnerID != userID {
		return fmt.Errorf("unauthorized: you don't own this project - Delete: ")
	}

	return s.repo.Delete(projectID)
}

func (s *service) AddUser(projectID int64, userID uuid.UUID, targetID uuid.UUID) error {
	project, err := s.repo.FindByID(projectID)
	if err != nil {
		return fmt.Errorf("AddUser: %v", err)
	}

	if project.OwnerID != userID {
		return fmt.Errorf("unauthorized: you don't own this project - AddUser: ")
	}

	projectMember, err := s.repo.MemberByID(projectID, targetID)
	if err != nil {
		return fmt.Errorf("AddUser: %v", err)
	}

	if projectMember != nil {
		if userID == projectMember.UserID {
			return fmt.Errorf("member already added to project - AddUser: %v", err)
		}
	}

	err = s.repo.AddUser(projectID, targetID)
	if err != nil {
		return fmt.Errorf("AddUser: %v", err)
	}

	return nil
}
