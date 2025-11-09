package task

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindByID(id int64) (*Task, error) {
	var task Task
	err := r.db.Where("id = ?", id).First(&task).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("task not found - FindByID: %v", err)
		}
		return nil, err
	}
	return &task, nil
}

func (r *Repository) FindByProjectID(projectID int64, filters map[string]interface{}) ([]*Task, error) {
	var tasks []*Task
	query := r.db.Where("project_id = ?", projectID)

	if status, ok := filters["status"]; ok {
		query = query.Where("status = ?", status)
	}
	if priority, ok := filters["priority"]; ok {
		query = query.Where("priority = ?", priority)
	}

	err := query.Order("created_at DESC").Find(&tasks).Error
	return tasks, fmt.Errorf("FindByProjectID: %v", err)
}

func (r *Repository) Create(task *Task) error {
	return fmt.Errorf("task - Create: %v", r.db.Create(task).Error)
}

func (r *Repository) Update(task *Task) error {
	return fmt.Errorf("task - Update: %v", r.db.Save(task).Error)
}

func (r *Repository) Delete(id int64) error {
	return fmt.Errorf("task - Delete: %v", r.db.Delete(&Task{}, id).Error)
}
