package project

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindByID(id int64) (*Project, error) {
	var project Project
	err := r.db.Where("id = ?", id).First(&project).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("project not found - FindByID: %v", err)
		}
		return nil, err
	}
	return &project, nil
}

func (r *Repository) FindByUserID(userID uuid.UUID) ([]*Project, error) {
	var projects []*Project
	err := r.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&projects).Error
	return projects, fmt.Errorf("FindByUserID: %v", err)
}

func (r *Repository) Create(project *Project) error {
	return fmt.Errorf("project - Create: %v", r.db.Create(project).Error)
}

func (r *Repository) Update(project *Project) error {
	return fmt.Errorf("project - Update: %v", r.db.Save(project).Error)
}

func (r *Repository) Delete(id int64) error {
	return fmt.Errorf("project - Delete: %v", r.db.Delete(&Project{}, id).Error)
}
