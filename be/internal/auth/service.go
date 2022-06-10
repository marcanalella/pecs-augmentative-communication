package auth

import (
	"crypto/rand"
	"fmt"
	"net/http"

	"github.com/pecs/pecs-be/internal/errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	log "github.com/sirupsen/logrus"

	"github.com/pecs/pecs-be/internal/entity"

	validator "github.com/go-playground/validator/v10"

	)

//Service encapsulates the authentication logic
type Service interface {
	//authenticate a user using username and password and return a JWT if authentication succeds
	Login(email string, password string) (*TokenDetails, error)

	// SignUp register a user and store it in database
	SignUp(user entity.User) (entity.User, error)

	//refresh a token user and store it
	Refresh(r *http.Request) (*TokenDetails, error)

	Password(email string) error

	ResetPassword(token string, email string, password string) error

	Validate(data interface{}) error
}

// Identity represents an authenticated user identity
type Identity interface {
	// GetID returns the user ID.
	GetID() uuid.UUID

	// GetName returns the user name.
	GetName() string

	GetRole() entity.Role
}

type service struct {
	repo                  Repository
	validator             *validator.Validate
	signinAccessTokenKey  string
	signinRefreshTokenKey string
	tokenExpiration       int
}

// NewService creates a new authentication service.
func NewService(repo Repository, validator *validator.Validate, signinAccessTokenKey string, signinRefreshTokenKey string, tokenExpiration int) Service {
	return service{repo,validator, signinAccessTokenKey, signinRefreshTokenKey, tokenExpiration}
}

func (s service) Login(email string, password string) (*TokenDetails, error) {
	identity, err := s.checkCredentials(email, password)

	if identity != nil {
		return s.generateJWT(identity)
	}

	return &TokenDetails{}, err
}

func (s service) SignUp(user entity.User) (entity.User, error) {
	identity, err := s.storeCredentials(user)

	if err != nil {
		return entity.User{}, err
	}

	return identity, nil
}

//todo must return a error not a string
func (s service) checkCredentials(email string, password string) (Identity, error) {
	user, err := s.repo.GetByEmail(email)

	if err != nil {
		return nil, errors.NewHTTPError(err, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), http.StatusText(http.StatusUnauthorized), "Authentication failed")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.NewHTTPError(err, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), http.StatusText(http.StatusUnauthorized), "Authentication failed")
	}

	log.Infof("Authentication succeeds")

	return entity.User{
		Base: entity.Base{
			ID: user.ID,
		},
		Name: user.Name,
		Role: user.Role,
	}, nil
}

func (s service) storeCredentials(user entity.User) (entity.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return entity.User{}, errors.NewHTTPError(err, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), http.StatusText(http.StatusInternalServerError), "User with that email already exists")
	}

	checkUser, _ := s.repo.GetByEmail(user.Email)

	if checkUser.Email != "" {
		return entity.User{}, errors.NewHTTPError(err, http.StatusConflict, http.StatusText(http.StatusConflict), http.StatusText(http.StatusConflict), "User with that email already exists")
	}

	user.ID = uuid.New()
	user.Password = string(hashedPassword)
	user.Role = "CUSTOMER" //CUSTOMER

	user, err = s.repo.Insert(user)

	if err != nil {
		return entity.User{}, errors.NewHTTPError(err, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), http.StatusText(http.StatusInternalServerError), "Issue with inserting into the database")
	}

	log.Info("Signup succeeds")

	return user, nil
}

func tokenGenerator() string {
	b := make([]byte, 10)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

type bodyMail struct {
	Email      string
	ResetToken string
}

func (s service) Password(email string) error {
	user, err := s.repo.GetByEmail(email)

	//TODO rimuovere in quanto anche se l'utente non esiste, dire comunque che la mail è stata inviata per il recupero password
	if err != nil {
		return errors.NewHTTPError(err, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), http.StatusText(http.StatusInternalServerError), "User with this email does not exist")
	}

	resetToken := tokenGenerator()

	user.ResetToken = resetToken
	user, err = s.repo.Update(user)

	if err != nil {
		return errors.NewHTTPError(err, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), http.StatusText(http.StatusInternalServerError), "Failed to set reset token")
	}

	return err
}

// ResetPassword check if reset_token for user is present in DB and equals to reset_token in body request
// set new password for this user
func (s service) ResetPassword(token string, email string, password string) error {
	//TODO GET USER EMAIL FROM TOKEN AND UPDATE PASSWORD
	user, err := s.repo.GetByEmail(email)

	//TODO rimuovere in quanto anche se l'utente non esiste, dire comunque che la mail è stata inviata per il recupero password
	if err != nil {
		return errors.NewHTTPError(err, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), http.StatusText(http.StatusInternalServerError), "User with this email does not exist")
	}

	//TODO SET AN EXPIRE DATE FOR TOKEN
	if user.ResetToken != token {
		return errors.NewHTTPError(err, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), http.StatusText(http.StatusBadRequest), "Reset token is invalid")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return errors.NewHTTPError(err, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), http.StatusText(http.StatusInternalServerError), "Failed to hash password")
	}

	user.Password = string(hashedPassword)
	user.ResetToken = ""
	user, err = s.repo.Update(user)

	if err != nil {
		return errors.NewHTTPError(err, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), http.StatusText(http.StatusBadRequest), "Failed to update pwd")
	}

	return nil
}

func (s service) Refresh(r *http.Request) (*TokenDetails, error) {
	return s.RefreshToken(r)
}

func (s service) Validate(data interface{}) error {
	err := s.validator.Struct(data)

	if err != nil {
		return errors.NewHTTPError(err, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), http.StatusText(http.StatusBadRequest), err.Error())
	}
	return nil
}
