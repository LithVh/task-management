package task

import (
	"fmt"
	"time"

	"task-management/internal/project"

	"github.com/google/uuid"
)

type Service interface {
	List(projectID int64, userID uuid.UUID, filters map[string]interface{}) ([]*TaskResponse, error)
	Create(projectID int64, userID uuid.UUID, dto *CreateTaskRequest) (*TaskResponse, error)
	GetByID(taskID int64, userID uuid.UUID) (*TaskResponse, error)
	Update(taskID int64, userID uuid.UUID, dto *UpdateTaskRequest) (*TaskResponse, error)
	ToggleComplete(taskID int64, userID uuid.UUID) (*TaskResponse, error)
	Delete(taskID int64, userID uuid.UUID) error
}

type service struct {
	repo        *Repository
	projectRepo *project.Repository
}

func NewService(repo *Repository, projectRepo *project.Repository) Service {
	return &service{
		repo:        repo,
		projectRepo: projectRepo,
	}
}

func (s *service) verifyProjectOwnership(projectID int64, userID uuid.UUID) error {
	proj, err := s.projectRepo.FindByID(projectID)
	if err != nil {
		return err
	}
	if proj.UserID != userID {
		return fmt.Errorf("verifyProjectOwnership: unauthorized: you don't own this project")
	}
	return nil
}

func (s *service) List(projectID int64, userID uuid.UUID, filters map[string]interface{}) ([]*TaskResponse, error) {
	// Verify project ownership
	if err := s.verifyProjectOwnership(projectID, userID); err != nil {
		return nil, err
	}

	tasks, err := s.repo.FindByProjectID(projectID, filters)
	if err != nil {
		return nil, err
	}
	return ToTaskResponseList(tasks), nil
}

func (s *service) Create(projectID int64, userID uuid.UUID, dto *CreateTaskRequest) (*TaskResponse, error) {
	// Verify project ownership
	if err := s.verifyProjectOwnership(projectID, userID); err != nil {
		return nil, err
	}

	now := time.Now()
	task := &Task{
		ProjectID:   projectID,
		Title:       dto.Title,
		Description: dto.Description,
		Status:      dto.Status,
		Priority:    dto.Priority,
		DueDate:     dto.DueDate,
		Completed:   dto.Status == StatusCompleted,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := s.repo.Create(task); err != nil {
		return nil, err
	}

	return ToTaskResponse(task), nil
}

func (s *service) GetByID(taskID int64, userID uuid.UUID) (*TaskResponse, error) {
	task, err := s.repo.FindByID(taskID)
	if err != nil {
		return nil, err
	}

	// Verify project ownership
	if err := s.verifyProjectOwnership(task.ProjectID, userID); err != nil {
		return nil, err
	}

	return ToTaskResponse(task), nil
}

func (s *service) Update(taskID int64, userID uuid.UUID, dto *UpdateTaskRequest) (*TaskResponse, error) {
	task, err := s.repo.FindByID(taskID)
	if err != nil {
		return nil, err
	}

	// Verify project ownership
	if err := s.verifyProjectOwnership(task.ProjectID, userID); err != nil {
		return nil, err
	}

	// Update fields if provided
	if dto.Title != "" {
		task.Title = dto.Title
	}
	if dto.Description != nil {
		task.Description = dto.Description
	}
	if dto.Status != "" {
		task.Status = dto.Status
		// Auto-update completed flag based on status
		task.Completed = dto.Status == StatusCompleted
	}
	if dto.Priority != nil {
		task.Priority = dto.Priority
	}
	if dto.DueDate != nil {
		task.DueDate = dto.DueDate
	}
	task.UpdatedAt = time.Now()

	if err := s.repo.Update(task); err != nil {
		return nil, err
	}

	return ToTaskResponse(task), nil
}

func (s *service) ToggleComplete(taskID int64, userID uuid.UUID) (*TaskResponse, error) {
	task, err := s.repo.FindByID(taskID)
	if err != nil {
		return nil, err
	}

	// Verify project ownership
	if err := s.verifyProjectOwnership(task.ProjectID, userID); err != nil {
		return nil, err
	}

	// Toggle completed
	task.Completed = !task.Completed

	// Update status accordingly
	if task.Completed {
		task.Status = StatusCompleted
	} else {
		task.Status = StatusTodo
	}

	task.UpdatedAt = time.Now()

	if err := s.repo.Update(task); err != nil {
		return nil, err
	}

	return ToTaskResponse(task), nil
}

func (s *service) Delete(taskID int64, userID uuid.UUID) error {
	task, err := s.repo.FindByID(taskID)
	if err != nil {
		return err
	}

	// Verify project ownership
	if err := s.verifyProjectOwnership(task.ProjectID, userID); err != nil {
		return err
	}

	return s.repo.Delete(taskID)
}
