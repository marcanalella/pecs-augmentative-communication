// Package version of Version API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
//     Schemes: http
//     Host: localhost
//     BasePath: /
//	   Version: 1.0.0
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//
// swagger:meta
package version

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pecs/pecs-be/internal/entity"
	log "github.com/sirupsen/logrus"
	"net/http"
	"sync/atomic"
	"time"
)

func RegisterHandlers(router *mux.Router, number string, build string) {

	isReady := &atomic.Value{}
	isReady.Store(false)
	go func() {
		log.Info("Readyz probe is negative by default ")
		time.Sleep(10 * time.Second)
		isReady.Store(true)
		log.Info("RReadyz probe is positive")

	}()

	router.HandleFunc("/version", getVersion(number, build)).Methods(http.MethodGet)
	router.HandleFunc("/healthz", healthz()).Methods(http.MethodGet)
	router.HandleFunc("/readyz", readyz(isReady))
}

// swagger:route GET /version version appVersion
// Returns version number and build time
// responses:
// 	200: Version
func getVersion(number string, build string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		version := &entity.Version{
			Number: number,
			Build:  build,
		}

		w.Header().Set("Content-Type", "application/json")

		err := json.NewEncoder(w).Encode(version)

		if err != nil {
			log.Error("Error encoding json version")
		}
	}
}

func healthz() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}

func readyz(isReady *atomic.Value) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		if isReady == nil || !isReady.Load().(bool) {
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
