package categories

import (
	"github.com/google/uuid"
	"github.com/pecs/pecs-be/internal/entity"
	"github.com/jinzhu/gorm"
)

type Repository interface {
	GetAllCategories(userId uuid.UUID) ([]entity.Category, error)

	InsertCategory(product entity.Category) (entity.Category, error)

	DeleteCategoryById(id uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new album repository
func NewRepository(db *gorm.DB) Repository {
	return repository{db}
}

func (r repository) GetAllCategories(userId uuid.UUID) ([]entity.Category, error) {
	var product []entity.Category
	err := r.db.Find(&product).Where("user_id = ?", userId).Error
	return product, err
}

func (r repository) InsertCategory(category entity.Category) (entity.Category, error) {
	err := r.db.Create(&category).Error
	return category, err
}

func (r repository) DeleteCategoryById(id uuid.UUID) error {
	var category entity.Category
	err := r.db.Delete(&category, "id = ?", id).Error
	return err
}
