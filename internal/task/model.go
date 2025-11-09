package task

import (
	"time"
)

type Task struct {
	ID          int64      `gorm:"primaryKey;autoIncrement"`
	ProjectID   int64      `gorm:"not null"`
	Title       string     `gorm:"type:varchar(255);not null"`
	Description *string    `gorm:"type:text"`
	Status      string     `gorm:"type:varchar(20);not null"`
	Priority    *string    `gorm:"type:varchar(20)"`
	DueDate     *time.Time `gorm:"type:timestamp"`
	Completed   bool       `gorm:"default:false;not null"`
	CreatedAt   time.Time  `gorm:"type:timestamp;not null"`
	UpdatedAt   time.Time  `gorm:"type:timestamp;not null"`
}

func (Task) TableName() string {
	return "task"
}

// Valid status values
const (
	StatusTodo       = "todo"
	StatusInProgress = "in_progress"
	StatusCompleted  = "completed"
)

// Valid priority values
const (
	PriorityLow    = "low"
	PriorityMedium = "medium"
	PriorityHigh   = "high"
)
