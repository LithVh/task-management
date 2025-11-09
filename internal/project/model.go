package project

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID          int64      `gorm:"primaryKey;autoIncrement"`
	UserID      uuid.UUID  `gorm:"type:uuid;not null"`
	Name        string     `gorm:"type:varchar(255);not null"`
	Description string     `gorm:"type:varchar(255)"`
	CreatedAt   time.Time  `gorm:"type:timestamp;not null"`
	UpdatedAt   time.Time  `gorm:"type:timestamp;not null"`
}

func (Project) TableName() string {
	return "project"
}
