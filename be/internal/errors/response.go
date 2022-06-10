package errors

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Error interface {
	Error() string
	ResponseBody() ([]byte, error)
	ResponseHeaders() (int, map[string]string)
}

type HTTPError struct {
	Cause             error  `json:"-"`
	Status            int    `json:"status"`
	StatusDescription string `json:"statusDescription"`
	Code              string `json:"code"`
	Detail            string `json:"detail"`
}

func (e *HTTPError) Error() string {
	if e.Cause == nil {
		return e.Detail
	}

	return e.Detail + " : " + e.Cause.Error()
}

func (e *HTTPError) ResponseBody() ([]byte, error) {
	body, err := json.Marshal(e)

	if err != nil {
		return nil, fmt.Errorf("error while parsing response body: %v", err)
	}

	return body, nil
}

func (e *HTTPError) ResponseHeaders() (int, map[string]string) {
	return e.Status, map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}
}

func NewHTTPError(err error, status int, statusDescription string, code string, detail string) error {
	return &HTTPError{
		Cause:             err,
		Status:            status,
		StatusDescription: statusDescription,
		Code:              code,
		Detail:            detail,
	}
}

func PrintError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	//w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8100")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization,Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,PUT,DELETE,PATCH")

	_, ok := err.(Error)

	//loggare qui dentro l'errore è corretto ???
	if ok {
		log.Error(err.(Error).Error())

		statusCode, _ := err.(Error).ResponseHeaders()
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
