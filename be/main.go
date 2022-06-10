package main

import (
	"flag"
	"github.com/dgrijalva/jwt-go"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"github.com/pecs/pecs-be/internal/actions"
	"github.com/pecs/pecs-be/internal/auth"
	"github.com/pecs/pecs-be/internal/categories"
	"github.com/pecs/pecs-be/internal/entity"
	"github.com/pecs/pecs-be/internal/errors"
	"github.com/pecs/pecs-be/internal/version"
	"github.com/pecs/pecs-be/sqldb"
	"github.com/rs/cors"

	"github.com/pecs/pecs-be/internal/config"
	log "github.com/sirupsen/logrus"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"

	goError "errors"
	_ "github.com/lib/pq"
)

var (
	number string
	build  string
)

func main() {
	flag.Parse()

	config.SetLogConfig()

	cfg := config.Load()
	dbCfg := cfg.Database

	db, err := sqldb.ConnectToDB(dbCfg.Dialect(), dbCfg.ConnectionInfo())
	if err != nil {
		log.Fatalf("Error connecting to DB %s", err.Error())
	}

	db.LogMode(true)
	db.AutoMigrate(&entity.User{},
		&entity.Category{},
		&entity.Action{},
	)
	db.Model(&entity.Category{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(&entity.Action{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT").AddForeignKey("category_id", "categories(id)", "RESTRICT", "RESTRICT")
	defer db.Close()

	log.Info(cfg)
	server := &http.Server{
		Addr:    cfg.Address + ":" + cfg.Port,
		Handler: buildHandler(db),
	}

	log.Info("Listening ", server.Addr)

	err = server.ListenAndServe()
	log.Fatalln(err)
}

// buildHandler sets up the HTTP routing and builds an HTTP handler.
func buildHandler(db *gorm.DB) http.Handler {
	entityValidator := validator.New()

	//all APIs are under "/api/v1" path prefix

	router := mux.NewRouter()
	router.Use(loggingMiddleware)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization,Content-Type"},
		AllowedMethods:   []string{"GET,POST,PUT,DELETE,PATCH,OPTIONS"},
	})

	// all APIs under noAuthGroup does not need that user is authenticated
	noAuthGroup := router.PathPrefix("/api/v1").Subrouter()

	// all APIs under authGroup need that user is authenticated
	authGroup := router.PathPrefix("/api/v1").Subrouter()
	authGroup.Use(tokenValidMiddleware)
	//authGroup.Use(hasAccess)

	auth.RegisterHandlers(noAuthGroup, auth.NewService(auth.NewRepository(db), entityValidator, "myatkey", "myrtkey", 1000000000000000)) //TODO this should be in env file
	version.RegisterHandlers(noAuthGroup, number, build)
	categories.RegisterHandlers(authGroup, categories.NewService(categories.NewRepository(db), entityValidator))
	actions.RegisterHandlers(authGroup, actions.NewService(actions.NewRepository(db), entityValidator))

	// all APIs under adminOnlyGroup need that user is authenticated and his role is "ADMIN"
	adminOnlyGroup := router.PathPrefix("/api/v1").Subrouter()
	adminOnlyGroup.Use(tokenValidMiddleware)

	// handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	router.Handle("/docs", sh).Methods(http.MethodGet)
	router.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	router.Handle("/images/{rest}", http.StripPrefix("/images/", http.FileServer(http.Dir("/resources/images"))))
	router.Handle("/images/{rest}/{rest}", http.StripPrefix("/images/", http.FileServer(http.Dir("/resources/images"))))

	handler := c.Handler(router)
	return handler
}

//Example function middleware, TODO remove
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("Service request to: " + r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func tokenValidMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := auth.TokenValid(r)

		if err != nil {
			errors.PrintError(w, err)
			return
		}

		log.Info("token validation passed")

		tokenString := token.Raw

		claims := jwt.MapClaims{}

		_, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("myatkey"), nil
		})

		if err != nil {
			log.Error("Error decoding JWT")
			log.Error(err.Error())
			return
		}

		userId := claims["id"]
		userRole := claims["role"]
		context.Set(r, "user_id", userId)
		context.Set(r, "user_role", userRole)

		// Call the next
		//handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

// authorize a user with Role!=Admin to only see his own data
// user with Role=Admin can see all data
func hasAccess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//TODO REFACTOR THIS
		// get ID from JWT
		tokenString := r.Header.Get("Authorization")

		claims := jwt.MapClaims{}

		_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("myatkey"), nil
		})

		if err != nil {
			log.Error("Error decoding JWT")
			return
		}

		jwtUserID := claims["id"]

		// get user_id from REST API path and check if it is equals to ID in JWT
		vars := mux.Vars(r)
		userID := vars["user_id"]

		if userID == jwtUserID {
			log.Info("CIAOOOO") //TODO CHANGE ASD TO ID PRESENT IN JWT
			next.ServeHTTP(w, r)
		} else {
			// customer can't access to this resource
			err := errors.NewHTTPError(goError.New(""), http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), http.StatusText(http.StatusUnauthorized), "Unauthorized to access to this resource")
			errors.PrintError(w, err)
			return
		}
	})
}
