package categories

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/pecs/pecs-be/internal/entity"
	"github.com/pecs/pecs-be/internal/errors"
	//log "github.com/sirupsen/logrus"
)

type Service interface {
	GetAllCategories(userId uuid.UUID) ([]entity.Category, error)

	InsertCategory(category entity.Category) (entity.Category, error)

	DeleteCategory(id uuid.UUID) error

	Validate(product *entity.Category) error
}

type service struct {
	repo      Repository
	validator *validator.Validate
}

// NewService creates a new authentication service.
func NewService(repo Repository, validator *validator.Validate) Service {
	return service{repo, validator}
}

func (s service) GetAllCategories(userId uuid.UUID) ([]entity.Category, error) {

	products, err := s.repo.GetAllCategories(userId)
	if err != nil {
		log.Error("Error on getting all category: ", err.Error())
		return nil, errors.NewHTTPError(err, http.StatusNotFound, http.StatusText(http.StatusNotFound), http.StatusText(http.StatusNotFound), "Error on getting all categories")
	}

	return products, nil
}

func (s service) InsertCategory(category entity.Category) (entity.Category, error) {
	category, err := s.repo.InsertCategory(category)
	if err != nil {
		log.Error("Error on inserting category: ", err.Error())
		return entity.Category{}, err
	}
	return category, nil
}

func (s service) DeleteCategory(id uuid.UUID) error {
	err := s.repo.DeleteCategoryById(id)
	if err != nil {
		log.Error("Error on deleting category", err.Error())
		return errors.NewHTTPError(err, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), http.StatusText(http.StatusInternalServerError), "Error on deleting product")
	}
	return nil
}

func (s service) Validate(product *entity.Category) error {
	err := s.validator.Struct(product)

	if err != nil {
		return errors.NewHTTPError(err, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), http.StatusText(http.StatusBadRequest), err.Error())
	}
	return nil
}
