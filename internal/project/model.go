package project

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID          int64      `gorm:"primaryKey;autoIncrement"`
	OwnerID     uuid.UUID  `gorm:"column:Owner_id;type:uuid;not null"`
	Name        string     `gorm:"type:varchar(255);not null"`
	Description string     `gorm:"type:varchar(255)"`
	CreatedAt   time.Time  `gorm:"type:timestamp;not null"`
	UpdatedAt   time.Time  `gorm:"type:timestamp;not null"`
}

func (Project) TableName() string {
	return "project"
}

type ProjectMember struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	ProjectID int64     `gorm:"column:Project_id;not null"`
	UserID    uuid.UUID `gorm:"column:User_id;type:uuid;not null"`
}

func (ProjectMember) TableName() string {
	return "project_member"
}
