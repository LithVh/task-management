package task

import (
	"context"
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

func (r *Repository) FindByID(ctx context.Context, id int64) (*Task, error) {
	var task Task
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&task).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("task not found - FindByID: %v", err)
		}
		return nil, err
	}
	return &task, nil
}

func (r *Repository) FindByProjectID(ctx context.Context, projectID int64, filters map[string]interface{}) ([]*Task, error) {
	var tasks []*Task
	query := r.db.WithContext(ctx).Where("project_id = ?", projectID)

	if status, ok := filters["status"]; ok {
		query = query.Where("status = ?", status)
	}
	if priority, ok := filters["priority"]; ok {
		query = query.Where("priority = ?", priority)
	}

	err := query.Order("created_at DESC").Find(&tasks).Error
	if err != nil {
		return nil, fmt.Errorf("FindByProjectID: %v", err)
	}
	return tasks, nil
}

func (r *Repository) Create(ctx context.Context, task *Task) error {
	err := r.db.WithContext(ctx).Create(task).Error
	if err != nil {
		return fmt.Errorf("task - Create: %v", err)
	}
	return nil
}

func (r *Repository) Update(ctx context.Context, task *Task) error {
	err := r.db.WithContext(ctx).Save(task).Error
	if err != nil {
		return fmt.Errorf("task - Update: %v", err)
	}
	return nil
}

func (r *Repository) Delete(ctx context.Context, id int64) error {
	err := r.db.WithContext(ctx).Delete(&Task{}, id).Error
	if err != nil {
		return fmt.Errorf("task - Delete: %v", err)
	}
	return nil
}
