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
	err := r.db.Where("owner_id = ?", userID).Order("created_at DESC").Find(&projects).Error
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

func (r *Repository) HasAccess(projectID int64, userID uuid.UUID) (bool, error) {
	// var project Project
	project, err := r.FindByID(projectID) 
	if err != nil {
		return false, fmt.Errorf("project doesnt exist - HasAccess: %v", err)
	}

	if project.OwnerID == userID {
		return true, nil // User is the owner
	}

	// err := r.db.Where("id = ? AND owner_id = ?", projectID, userID).First(&project).Error
	// if err == nil {
	// 	return true, nil // User is the owner
	// }
	// if !errors.Is(err, gorm.ErrRecordNotFound) {
	// 	return false, fmt.Errorf("owner check - HasAccess: %v", err)
	// }

	// Check if user is a member
	var member ProjectMember
	err = r.db.Where("project_id = ? AND user_id = ?", projectID, userID).First(&member).Error
	if err == nil {
		return true, nil // User is a member
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil // User has no access
	}
	return false, fmt.Errorf("HasAccess: %v", err)
}
