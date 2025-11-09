package task

import (
	"time"
)

type CreateTaskRequest struct {
	Title       string     `json:"title" binding:"required,min=2,max=255"`
	Description *string    `json:"description" binding:"omitempty"`
	Status      string     `json:"status" binding:"required,oneof=todo in_progress completed"`
	Priority    *string    `json:"priority" binding:"omitempty,oneof=low medium high"`
	DueDate     *time.Time `json:"due_date" binding:"omitempty"`
}

type UpdateTaskRequest struct {
	Title       string     `json:"title" binding:"omitempty,min=2,max=255"`
	Description *string    `json:"description" binding:"omitempty"`
	Status      string     `json:"status" binding:"omitempty,oneof=todo in_progress completed"`
	Priority    *string    `json:"priority" binding:"omitempty,oneof=low medium high"`
	DueDate     *time.Time `json:"due_date" binding:"omitempty"`
}

type TaskResponse struct {
	ID          int64      `json:"id"`
	ProjectID   int64      `json:"project_id"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Status      string     `json:"status"`
	Priority    *string    `json:"priority"`
	DueDate     *time.Time `json:"due_date"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func ToTaskResponse(task *Task) *TaskResponse {
	return &TaskResponse{
		ID:          task.ID,
		ProjectID:   task.ProjectID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		Priority:    task.Priority,
		DueDate:     task.DueDate,
		Completed:   task.Completed,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}

func ToTaskResponseList(tasks []*Task) []*TaskResponse {
	responses := make([]*TaskResponse, len(tasks))
	for i, task := range tasks {
		responses[i] = ToTaskResponse(task)
	}
	return responses
}
