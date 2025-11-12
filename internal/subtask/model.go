package subtask

import (
	"time"

	"github.com/google/uuid"
)

const (
	StatusTodo       = "todo"
	StatusInProgress = "in_progress"
	StatusCompleted  = "completed"

	PriorityLow    = "low"
	PriorityMedium = "medium"
	PriorityHigh   = "high"
)

type Subtask struct {
	ID          int64      `gorm:"primaryKey;autoIncrement"`
	TaskID      int64      `gorm:"not null;index"`
	AssignedTo  *uuid.UUID `gorm:"type:uuid"`
	Title       string     `gorm:"type:varchar(255);not null"`
	Description *string    `gorm:"type:text"`
	Status      string     `gorm:"type:varchar(20);not null;default:todo"`
	Priority    *string    `gorm:"type:varchar(20)"`
	DueDate     *time.Time `gorm:"type:timestamp"`
	Completed   bool       `gorm:"default:false;not null"`
	CreatedAt   time.Time  `gorm:"type:timestamp;not null"`
	UpdatedAt   time.Time  `gorm:"type:timestamp;not null"`
}

func (Subtask) TableName() string {
	return "subtask"
}
