// Package auth of Authentication API
//
// This should demonstrate all the possible comment annotations!
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
//     Schemes: http
//     Host: localhost
//     BasePath: /
//	   Version: 1.0.1
//     Consumes:
//     - application/json
//
//
//     Produces:
//     - application/json
//
// swagger:meta
package auth

import (
	"bytes"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/pecs/pecs-be/internal/entity"
	"github.com/pecs/pecs-be/internal/errors"

	"github.com/gorilla/mux"
)

func RegisterHandlers(router *mux.Router, service Service) {
	router.HandleFunc("/login", login(service)).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/signup", signup(service)).Methods(http.MethodPost)
	router.HandleFunc("/refresh", refresh(service)).Methods(http.MethodPost, http.MethodOptions)
	//router.HandleFunc("/logout", logout(service)).Methods(http.MethodPost) //TODO
	router.HandleFunc("/password/email", passwordEmail(service)).Methods(http.MethodPost)
	router.HandleFunc("/password/reset", passwordReset(service)).Methods(http.MethodPost)
}

type loginRequest struct {
	Email    string `json:"email" validate:"required,email,lowercase"`
	Password string `json:"password" validate:"required,max=128"`
}

type passwordRequest struct {
	Email string `json:"email" validate:"required,email,lowercase"`
}

type resetPasswordRequest struct {
	Token    string `json:"token" validate:"required"`
	Email    string `json:"email" validate:"required,email,lowercase"`
	Password string `json:"password" validate:"required,min=6,max=128"`
}

// swagger:route POST /login login userLogin
// Authenticate the user through user and password
// responses:
// 	200: LoginResponse
func login(service Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var lr loginRequest

		err := json.NewDecoder(r.Body).Decode(&lr)
		if err != nil {
			buf := new(bytes.Buffer)
			buf.ReadFrom(r.Body)
			errors.PrintError(w, errors.NewHTTPError(err, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), http.StatusText(http.StatusNotFound), "LOGIN - Error on decoding request"))
			return
		}
		log.Info("Auth:Login request decoded")


		err = service.Validate(&lr)
		if err != nil {
			log.Error("LOGIN - Error on validating request")
			errors.PrintError(w, err)
			return
		}
		log.Info("Auth:Login request validated for " + lr.Email)

		token, err := service.Login(lr.Email, lr.Password)
		if err != nil {
			errors.PrintError(w, err)
			return
		}

		//set refresh token on http cookie only
		//refreshToken := &http.Cookie{Name: "refresh_token", Domain: "http://localhost:8100", Value: token.RefreshToken, Path: "/", HttpOnly: true, Secure: false, SameSite: http.SameSiteNoneMode}
		//http.SetCookie(w, refreshToken)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err = json.NewEncoder(w).Encode(token)

		if err != nil {
			errors.PrintError(w, err)
			return
		}
	}
}

// swagger:route POST /signup signup signupUser
// Signup user given the request body
// responses:
// 	200: userResponse
func signup(service Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req entity.User

		err := json.NewDecoder(r.Body).Decode(&req)

		if err != nil {
			errors.PrintError(w, err)
			return
		}

		err = service.Validate(&req)

		if err != nil {
			errors.PrintError(w, err)
			return
		}

		user := entity.User{
			Email:           req.Email,
			Password:        req.Password,
			Surname:         req.Surname,
			Name:            req.Name,
		}

		user, err = service.SignUp(user)

		if err != nil {
			errors.PrintError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}

// swagger:route POST /logout logout userLogin
// Logout the user from the website
// responses:
// 	201: LogoutResponse
/*func logout(service Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}*/

func refresh(service Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenDetails, err := service.Refresh(r)

		if err != nil {
			errors.PrintError(w, err)
			return
		}

		//set refresh token on http cookie only
		//refreshToken := &http.Cookie{Name: "refresh_token" /*Domain: "http://localhost:8100",*/, Value: tokenDetails.RefreshToken, Path: "/", HttpOnly: true, Secure: false, SameSite: http.SameSiteNoneMode}
		//http.SetCookie(w, refreshToken)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(tokenDetails)
	}
}

func passwordEmail(service Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var pr passwordRequest

		err := json.NewDecoder(r.Body).Decode(&pr)

		if err != nil {
			errors.PrintError(w, err)
			return
		}

		err = service.Validate(&pr)

		if err != nil {
			errors.PrintError(w, err)
			return
		}

		err = service.Password(pr.Email)

		if err != nil {
			errors.PrintError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func passwordReset(service Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var rpr resetPasswordRequest

		err := json.NewDecoder(r.Body).Decode(&rpr)

		if err != nil {
			errors.PrintError(w, err)
			return
		}

		err = service.Validate(&rpr)

		if err != nil {
			errors.PrintError(w, err)
			return
		}

		err = service.ResetPassword(rpr.Token, rpr.Email, rpr.Password)

		if err != nil {
			errors.PrintError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
	}
}
