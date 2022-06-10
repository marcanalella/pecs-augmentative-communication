package actions

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/pecs/pecs-be/internal/entity"
)

type Repository interface {
	GetAllActionsByCategoryId(userId uuid.UUID, categoryId uuid.UUID) ([]entity.Action, error)

	InsertAction(product entity.Action) (entity.Action, error)

	DeleteActionById(id uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new album repository
func NewRepository(db *gorm.DB) Repository {
	return repository{db}
}

func (r repository) GetAllActionsByCategoryId(userId uuid.UUID, categoryId uuid.UUID) ([]entity.Action, error) {
	var action []entity.Action
	err := r.db.Where("user_id = ? and category_id = ?", userId, categoryId).Find(&action).Error
	return action, err
}

func (r repository) InsertAction(action entity.Action) (entity.Action, error) {
	err := r.db.Create(&action).Error
	return action, err
}

func (r repository) DeleteActionById(id uuid.UUID) error {
	var action entity.Action
	err := r.db.Delete(&action, "id = ?", id).Error
	return err
}
