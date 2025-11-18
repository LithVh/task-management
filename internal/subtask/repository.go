package subtask

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

func (r *Repository) FindByID(ctx context.Context, id int64) (*Subtask, error) {
	var subtask Subtask
	if err := r.db.WithContext(ctx).First(&subtask, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("subtask not found - FindByID")
		}
		return nil, fmt.Errorf("failed to find subtask - FindByID: %v", err)
	}
	return &subtask, nil
}

func (r *Repository) FindByTaskID(ctx context.Context, taskID int64, status, priority *string) ([]Subtask, error) {
	query := r.db.WithContext(ctx).Where("task_id = ?", taskID)

	//only display task with certain filters
	if status != nil && *status != "" {
		query = query.Where("status = ?", *status)
	}

	if priority != nil && *priority != "" {
		query = query.Where("priority = ?", *priority)
	}

	var subtasks []Subtask
	if err := query.Order("created_at DESC").Find(&subtasks).Error; err != nil {
		return nil, fmt.Errorf("failed to find subtasks - FindByTaskID: %v", err)
	}

	return subtasks, nil
}

func (r *Repository) CountByTaskID(ctx context.Context, taskID int64, completed *bool) (int64, error) {
	query := r.db.WithContext(ctx).Model(&Subtask{}).Where("task_id = ?", taskID)

	if completed != nil {
		query = query.Where("completed = ?", *completed)
	}

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count subtasks - CountByTaskID: %v", err)
	}

	return count, nil
}

func (r *Repository) Create(ctx context.Context, subtask *Subtask) error {
	if err := r.db.WithContext(ctx).Create(subtask).Error; err != nil {
		return fmt.Errorf("failed to create subtask - Create: %v", err)
	}
	return nil
}

func (r *Repository) Update(ctx context.Context, subtask *Subtask) error {
	if err := r.db.WithContext(ctx).Save(subtask).Error; err != nil {
		return fmt.Errorf("failed to update subtask - Update: %v", err)
	}
	return nil
}

func (r *Repository) Delete(ctx context.Context, id int64) error {
	if err := r.db.WithContext(ctx).Delete(&Subtask{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete subtask - Delete: %v", err)
	}
	return nil
}
