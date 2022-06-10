package actions

import (
	"encoding/json"
	"fmt"
	"github.com/pecs/pecs-be/internal/entity"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/pecs/pecs-be/internal/errors"
)

func RegisterHandlers(router *mux.Router, service Service) {
	router.HandleFunc("/actions/{categoryId}", getAllActionsByCategoryId(service)).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/action/{id}", deleteAction(service)).Methods(http.MethodDelete)
	router.HandleFunc("/action", insertAction(service)).Methods(http.MethodPost, http.MethodOptions)
}

func getAllActionsByCategoryId(service Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		categoryID, err := uuid.Parse(vars["categoryId"])
		if err != nil {
			log.Error("Error on parsing category id", err.Error())
			errors.PrintError(w, errors.NewHTTPError(err, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), http.StatusText(http.StatusBadRequest), "Error on parsing id"))
			return
		}

		log.Info(context.Get(r, "user_id"))
		userId, err := uuid.Parse(fmt.Sprint(context.Get(r, "user_id")))
		if err != nil {
			log.Error("Error on parsing uuid: ", err.Error())
			errors.PrintError(w, err)
			return
		}

		product, err := service.GetAllActionsByCategoryId(userId, categoryID)
		if err != nil {
			errors.PrintError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(product)

		if err != nil {
			errors.PrintError(w, err)
			return
		}
	}
}

func insertAction(service Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req entity.Action

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Error("Error decoding action request: ", err.Error())
			errors.PrintError(w, err)
			return
		}

		err = service.Validate(&req)
		if err != nil {
			log.Error("Error validating action: ", err.Error())
			errors.PrintError(w, err)
			return
		}

		log.Info(context.Get(r, "user_id"))
		userId, err := uuid.Parse(fmt.Sprint(context.Get(r, "user_id")))
		if err != nil {
			log.Error("Error on parsing uuid: ", err.Error())
			errors.PrintError(w, err)
			return
		}
		req.UserId = userId
		log.Info("User ID inside request")

		action, err := service.InsertAction(req)
		if err != nil {
			log.Error("Error on inserting action: ", err.Error())
			errors.PrintError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		err = json.NewEncoder(w).Encode(action)
		if err != nil {
			errors.PrintError(w, err)
			return
		}
	}
}

func deleteAction(service Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		actionID, err := uuid.Parse(vars["id"])
		if err != nil {
			log.Error("Error on parsing action id", err.Error())
			errors.PrintError(w, errors.NewHTTPError(err, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), http.StatusText(http.StatusBadRequest), "Error on parsing id"))
			return
		}

		err = service.DeleteAction(actionID)
		if err != nil {
			errors.PrintError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}
