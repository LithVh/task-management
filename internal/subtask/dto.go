package subtask

import (
	"time"

	"github.com/google/uuid"
)

type CreateSubtaskRequest struct {
	Title       string     `json:"title" binding:"required,min=2,max=255"`
	Description *string    `json:"description"`
	Status      string     `json:"status" binding:"required,oneof=todo in_progress completed"`
	Priority    *string    `json:"priority" binding:"omitempty,oneof=low medium high"`
	DueDate     *time.Time `json:"due_date"`
	AssignedTo  *uuid.UUID `json:"assigned_to"`
}

type UpdateSubtaskRequest struct {
	Title       *string    `json:"title" binding:"omitempty,min=2,max=255"`
	Description *string    `json:"description"`
	Status      *string    `json:"status" binding:"omitempty,oneof=todo in_progress completed"`
	Priority    *string    `json:"priority" binding:"omitempty,oneof=low medium high"`
	DueDate     *time.Time `json:"due_date"`
	AssignedTo  *uuid.UUID `json:"assigned_to"`
}

type SubtaskResponse struct {
	ID          int64      `json:"id"`
	TaskID      int64      `json:"task_id"`
	AssignedTo  *uuid.UUID `json:"assigned_to"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Status      string     `json:"status"`
	Priority    *string    `json:"priority"`
	DueDate     *time.Time `json:"due_date"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func ToSubtaskResponse(subtask *Subtask) SubtaskResponse {
	return SubtaskResponse{
		ID:          subtask.ID,
		TaskID:      subtask.TaskID,
		AssignedTo:  subtask.AssignedTo,
		Title:       subtask.Title,
		Description: subtask.Description,
		Status:      subtask.Status,
		Priority:    subtask.Priority,
		DueDate:     subtask.DueDate,
		Completed:   subtask.Completed,
		CreatedAt:   subtask.CreatedAt,
		UpdatedAt:   subtask.UpdatedAt,
	}
}

func ToSubtaskResponseList(subtasks []Subtask) []SubtaskResponse {
	responses := make([]SubtaskResponse, len(subtasks))
	for i, subtask := range subtasks {
		responses[i] = ToSubtaskResponse(&subtask)
	}
	return responses
}
