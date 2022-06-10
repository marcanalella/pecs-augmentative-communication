package auth

import (
	"github.com/google/uuid"
	"github.com/pecs/pecs-be/internal/entity"
	"github.com/jinzhu/gorm"
)

// Repository encapsulates the logic to access albums from the data source.
type Repository interface {
	GetByEmail(email string) (entity.User, error)

	Insert(user entity.User) (entity.User, error)

	Update(user entity.User) (entity.User, error)

	GetById(id uuid.UUID) (entity.User, error)
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new album repository
func NewRepository(db *gorm.DB) Repository {
	return repository{db}
}

// Get reads the album with the specified ID from the database.
func (r repository) GetByEmail(email string) (entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (r repository) GetById(id uuid.UUID) (entity.User, error) {
	var user entity.User
	err := r.db.Where("id = ?", id).First(&user).Error
	return user, err
}

func (r repository) Insert(user entity.User) (entity.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r repository) Update(user entity.User) (entity.User, error) {
	err := r.db.Where("id = ?", user.ID).Save(user).Error
	return user, err
}
