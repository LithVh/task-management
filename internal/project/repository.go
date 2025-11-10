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
	if err != nil {
		return nil, fmt.Errorf("FindByUserID: %v", err)
	}
	return projects, nil
}

func (r *Repository) Create(project *Project) error {
	err := r.db.Create(project).Error
	if err != nil {
		return fmt.Errorf("project - Create: %v", err)
	}
	return nil
}

func (r *Repository) Update(project *Project) error {
	err := r.db.Save(project).Error
	if err != nil {
		return fmt.Errorf("project - Update: %v", err)
	}
	return nil
}

func (r *Repository) Delete(id int64) error {
	err := r.db.Delete(&Project{}, id).Error
	if err != nil {
		return fmt.Errorf("project - Delete: %v", err)
	}
	return nil
}
