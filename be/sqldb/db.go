package sqldb

import (
	"github.com/jinzhu/gorm"
)

type DbHandler struct {
	db *gorm.DB
}

// NewDbHandler returns a new DbHandler
func NewDbHandler(db *gorm.DB) *DbHandler {
	return &DbHandler{
		db: db,
	}
}

//ConnectToDB used for creating connection pool to DB
func ConnectToDB(dialect string, connectionInfo string) (*gorm.DB, error) {
	return gorm.Open(dialect, connectionInfo)
}
