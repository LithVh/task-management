package user

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

func (r *Repository) FindByID(id uuid.UUID) (*User, error) {
	var user User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("findByID: %v", err)
	}

	return &user, nil
}

func (r *Repository) FindByEmail(email string) (*User, error) {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("FindByEmail: %v", err)
	}

	return &user, nil
}

func (r *Repository) Create(user *User) error {
	err := r.db.Create(user).Error
	if err != nil {
		return fmt.Errorf("Create: %v", err)
	}
	return nil
}

func (r *Repository) Update(user *User) error {
	err := r.db.Model(&user).Where("id = ?", user.ID).Updates(user).Error
	if err != nil {
		return fmt.Errorf("Update: %v", err)
	}
	return nil
}

func (r *Repository) EmailAvailale(email string) (bool, error) {
	var count int64
	err := r.db.Model(&User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, fmt.Errorf("EmailExists: %v", err)
	}
	return count < 1, nil
}
