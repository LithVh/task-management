package task

import (
	"context"
	"fmt"
	"time"

	"task-management/internal/project"

	"github.com/google/uuid"
)

type Service interface {
	List(ctx context.Context, projectID int64, userID uuid.UUID, filters map[string]interface{}) ([]*TaskResponse, error)
	Create(ctx context.Context, projectID int64, userID uuid.UUID, dto *CreateTaskRequest) (*TaskResponse, error)
	GetByID(ctx context.Context, taskID int64, userID uuid.UUID) (*TaskResponse, error)
	Update(ctx context.Context, taskID int64, userID uuid.UUID, dto *UpdateTaskRequest) (*TaskResponse, error)
	ToggleComplete(ctx context.Context, taskID int64, userID uuid.UUID) (*TaskResponse, error)
	Delete(ctx context.Context, taskID int64, userID uuid.UUID) error
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

func (s *service) verifyProjectOwner(ctx context.Context, projectID int64, userID uuid.UUID) error {
	hasAccess, err := s.projectRepo.IsOwner(ctx, projectID, userID)
	if err != nil {
		return fmt.Errorf("verifyProjectAccess: %v", err)
	}
	if !hasAccess {
		return fmt.Errorf("verifyProjectAccess: unauthorized: you are not the owner of this project")
	}
	return nil
}

func (s *service) verifyProjectAccess(ctx context.Context, projectID int64, userID uuid.UUID) error {
	hasAccess, err := s.projectRepo.HasAccess(ctx, projectID, userID)
	if err != nil {
		return fmt.Errorf("verifyProjectAccess: %v", err)
	}
	if !hasAccess {
		return fmt.Errorf("verifyProjectAccess: unauthorized: you don't have access to this project")
	}
	return nil
}

func (s *service) List(ctx context.Context, projectID int64, userID uuid.UUID, filters map[string]interface{}) ([]*TaskResponse, error) {
	if err := s.verifyProjectAccess(ctx, projectID, userID); err != nil {
		return nil, err
	}

	tasks, err := s.repo.FindByProjectID(ctx, projectID, filters)
	if err != nil {
		return nil, err
	}
	return ToTaskResponseList(tasks), nil
}

func (s *service) Create(ctx context.Context, projectID int64, userID uuid.UUID, dto *CreateTaskRequest) (*TaskResponse, error) {
	if err := s.verifyProjectAccess(ctx, projectID, userID); err != nil {
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

	if err := s.repo.Create(ctx, task); err != nil {
		return nil, err
	}

	return ToTaskResponse(task), nil
}

func (s *service) GetByID(ctx context.Context, taskID int64, userID uuid.UUID) (*TaskResponse, error) {
	task, err := s.repo.FindByID(ctx, taskID)
	if err != nil {
		return nil, fmt.Errorf("GetByID: %v", err)
	}

	if err := s.verifyProjectAccess(ctx, task.ProjectID, userID); err != nil {
		return nil, fmt.Errorf("GetByID: %v", err)
	}

	return ToTaskResponse(task), nil
}

func (s *service) Update(ctx context.Context, taskID int64, userID uuid.UUID, dto *UpdateTaskRequest) (*TaskResponse, error) {
	task, err := s.repo.FindByID(ctx, taskID)
	if err != nil {
		return nil, fmt.Errorf("Update: %v", err)
	}

	if err := s.verifyProjectOwner(ctx, task.ProjectID, userID); err != nil {
		return nil, fmt.Errorf("Update: %v", err)
	}

	if dto.Title != "" {
		task.Title = dto.Title
	}
	if dto.Description != nil {
		task.Description = dto.Description
	}
	if dto.Status != "" {
		task.Status = dto.Status
		task.Completed = dto.Status == StatusCompleted
	}
	if dto.Priority != nil {
		task.Priority = dto.Priority
	}
	if dto.DueDate != nil {
		task.DueDate = dto.DueDate
	}
	task.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, task); err != nil {
		return nil, fmt.Errorf("Update: %v", err)
	}

	return ToTaskResponse(task), nil
}

func (s *service) ToggleComplete(ctx context.Context, taskID int64, userID uuid.UUID) (*TaskResponse, error) {
	task, err := s.repo.FindByID(ctx, taskID)
	if err != nil {
		return nil, fmt.Errorf("ToggleComplete: %v", err)
	}

	if err := s.verifyProjectAccess(ctx, task.ProjectID, userID); err != nil {
		return nil, err
	}

	task.Completed = !task.Completed

	if task.Completed {
		task.Status = StatusCompleted
	} else {
		task.Status = StatusTodo
	}

	task.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, task); err != nil {
		return nil, err
	}

	return ToTaskResponse(task), nil
}

func (s *service) Delete(ctx context.Context, taskID int64, userID uuid.UUID) error {
	task, err := s.repo.FindByID(ctx, taskID)
	if err != nil {
		return err
	}

	if err := s.verifyProjectOwner(ctx, task.ProjectID, userID); err != nil {
		return err
	}

	return s.repo.Delete(ctx, taskID)
}
