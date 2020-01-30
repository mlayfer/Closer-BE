package dataaccess

import (
	"github.com/jinzhu/gorm"

	_  "github.com/jinzhu/gorm/dialects/sqlite"
)

type Database struct {
	db *gorm.DB
}

func NewDatabaseAccess() (dbConn *Database, closure func() error) {
	db, err := gorm.Open("sqlite3", ".\\CloserDB")
	if err != nil {
		panic("failed to open DB connection")
	}

	return &Database{db: db},db.Close
}