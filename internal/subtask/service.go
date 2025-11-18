package subtask

import (
	"context"
	"fmt"
	"strings"
	"time"

	"task-management/internal/task"

	"github.com/google/uuid"
)

type Service interface {
	List(ctx context.Context, userID uuid.UUID, taskID int64, status, priority *string) ([]Subtask, error)
	Create(ctx context.Context, userID uuid.UUID, taskID int64, req CreateSubtaskRequest) (*Subtask, error)
	GetByID(ctx context.Context, userID uuid.UUID, id int64) (*Subtask, error)
	Update(ctx context.Context, userID uuid.UUID, id int64, req UpdateSubtaskRequest) (*Subtask, error)
	ToggleComplete(ctx context.Context, userID uuid.UUID, id int64) (*Subtask, error)
	Delete(ctx context.Context, userID uuid.UUID, id int64) error
}

type service struct {
	repo        *Repository
	taskService task.Service
	// projectService project.Service
}

func NewService(repo *Repository, taskService task.Service) Service {
	return &service{
		repo:        repo,
		taskService: taskService,
	}
}

func (s *service) verifyProjectAccess(ctx context.Context, userID uuid.UUID, taskID int64) error {
	_, err := s.taskService.GetByID(ctx, taskID, userID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return fmt.Errorf("task not found")
		}
		if strings.Contains(err.Error(), "does not belong to user") {
			return fmt.Errorf("unauthorized: no access to this")
		}
		return err
	}
	return nil
}

func (s *service) checkAndAutoCompleteTask(ctx context.Context, taskID int64, userID uuid.UUID) error {
	totalCount, err := s.repo.CountByTaskID(ctx, taskID, nil)
	if err != nil {
		return err
	}

	if totalCount == 0 {
		return nil
	}

	completed := true
	completedCount, err := s.repo.CountByTaskID(ctx, taskID, &completed)
	if err != nil {
		return err
	}

	// If all subtasks are completed, mark parent task as completed
	if completedCount == totalCount {
		parentTask, err := s.taskService.GetByID(ctx, taskID, userID)
		if err != nil {
			return err
		}

		// Only auto-complete if not already completed
		if parentTask.Status != task.StatusCompleted {
			completedStatus := task.StatusCompleted
			_, err = s.taskService.Update(ctx, taskID, userID, &task.UpdateTaskRequest{
				Status: completedStatus,
			})
			if err != nil {
				return fmt.Errorf("failed to auto-complete parent task: %w", err)
			}
		}
	}

	return nil
}

func (s *service) List(ctx context.Context, userID uuid.UUID, taskID int64, status, priority *string) ([]Subtask, error) {
	// Verify user has access to the parent task
	if err := s.verifyProjectAccess(ctx, userID, taskID); err != nil {
		return nil, err
	}

	subtasks, err := s.repo.FindByTaskID(ctx, taskID, status, priority)
	if err != nil {
		return nil, err
	}

	return subtasks, nil
}

func (s *service) Create(ctx context.Context, userID uuid.UUID, taskID int64, req CreateSubtaskRequest) (*Subtask, error) {
	// Verify user has access to the parent task
	if err := s.verifyProjectAccess(ctx, userID, taskID); err != nil {
		return nil, err
	}

	now := time.Now()
	completed := req.Status == StatusCompleted

	subtask := &Subtask{
		TaskID:      taskID,
		AssignedTo:  req.AssignedTo,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		Priority:    req.Priority,
		DueDate:     req.DueDate,
		Completed:   completed,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := s.repo.Create(ctx, subtask); err != nil {
		return nil, err
	}

	return subtask, nil
}

func (s *service) GetByID(ctx context.Context, userID uuid.UUID, id int64) (*Subtask, error) {
	subtask, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := s.verifyProjectAccess(ctx, userID, subtask.TaskID); err != nil {
		return nil, err
	}

	return subtask, nil
}

func (s *service) Update(ctx context.Context, userID uuid.UUID, id int64, req UpdateSubtaskRequest) (*Subtask, error) {
	subtask, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("subtask not found - Update: %v", err)
	}

	if err := s.verifyProjectAccess(ctx, userID, subtask.TaskID); err != nil {
		return nil, err
	}

	// Update fields if provided
	if req.Title != nil && *req.Title != "" {
		subtask.Title = *req.Title
	}

	if req.Description != nil {
		subtask.Description = req.Description
	}

	if req.Status != nil && *req.Status != "" {
		subtask.Status = *req.Status
		subtask.Completed = *req.Status == StatusCompleted
	}

	if req.Priority != nil {
		subtask.Priority = req.Priority
	}

	if req.DueDate != nil {
		subtask.DueDate = req.DueDate
	}

	if req.AssignedTo != nil {
		subtask.AssignedTo = req.AssignedTo
	}

	subtask.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, subtask); err != nil {
		return nil, err
	}

	// Check if all subtasks are completed to auto-complete parent task
	if err := s.checkAndAutoCompleteTask(ctx, subtask.TaskID, userID); err != nil {
		// Log the error but don't fail the subtask update
		fmt.Printf("Warning: failed to auto-complete parent task: %v\n", err)
	}

	return subtask, nil
}

func (s *service) ToggleComplete(ctx context.Context, userID uuid.UUID, id int64) (*Subtask, error) {
	subtask, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := s.verifyProjectAccess(ctx, userID, subtask.TaskID); err != nil {
		return nil, err
	}

	// Toggle completion status
	subtask.Completed = !subtask.Completed
	if subtask.Completed {
		subtask.Status = StatusCompleted
	} else {
		subtask.Status = StatusInProgress
	}

	subtask.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, subtask); err != nil {
		return nil, err
	}

	// check if all subtasks are completed to auto-complete parent task
	if err := s.checkAndAutoCompleteTask(ctx, subtask.TaskID, userID); err != nil {
		//log the error but don't fail the subtask update
		fmt.Printf("Warning: failed to auto-complete parent task: %v\n", err)
	}

	return subtask, nil
}

func (s *service) Delete(ctx context.Context, userID uuid.UUID, id int64) error {
	subtask, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if err := s.verifyProjectAccess(ctx, userID, subtask.TaskID); err != nil {
		return err
	}

	return s.repo.Delete(ctx, id)
}
