package actions

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
	GetAllActionsByCategoryId(userId uuid.UUID, categoryId uuid.UUID) ([]entity.Action, error)

	InsertAction(action entity.Action) (entity.Action, error)

	DeleteAction(id uuid.UUID) error

	Validate(product *entity.Action) error
}

type service struct {
	repo      Repository
	validator *validator.Validate
}

// NewService creates a new authentication service.
func NewService(repo Repository, validator *validator.Validate) Service {
	return service{repo, validator}
}

func (s service) GetAllActionsByCategoryId(userId uuid.UUID, categoryId uuid.UUID) ([]entity.Action, error) {

	products, err := s.repo.GetAllActionsByCategoryId(userId, categoryId)
	if err != nil {
		log.Error("Error on getting all action: ", err.Error())
		return nil, errors.NewHTTPError(err, http.StatusNotFound, http.StatusText(http.StatusNotFound), http.StatusText(http.StatusNotFound), "Error on getting all categories")
	}

	return products, nil
}

func (s service) InsertAction(action entity.Action) (entity.Action, error) {
	action, err := s.repo.InsertAction(action)
	if err != nil {
		log.Error("Error on inserting action: ", err.Error())
		return entity.Action{}, err
	}
	return action, nil
}

func (s service) DeleteAction(id uuid.UUID) error {
	err := s.repo.DeleteActionById(id)
	if err != nil {
		log.Error("Error on deleting action", err.Error())
		return errors.NewHTTPError(err, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), http.StatusText(http.StatusNotFound), "Error on deleting product")
	}
	return nil
}

func (s service) Validate(product *entity.Action) error {
	err := s.validator.Struct(product)

	if err != nil {
		return errors.NewHTTPError(err, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), http.StatusText(http.StatusNotFound), err.Error())
	}
	return nil
}
