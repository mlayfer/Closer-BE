package dataaccess

import (
	"Closer/dataaccess/sqliteImpl"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Factory struct{}

func (f *Factory) CreateDataAccess(dbPath string) (DataAccess, func() error) {
	db, err := gorm.Open("sqlite3", dbPath)
	if err != nil {
		panic(fmt.Sprintf("failed to open DB connection with err %s", err.Error()))
	}

	return DataAccess{
		platform: sqliteImpl.NewPlatformDB(db),
		user:     sqliteImpl.NewUserDB(db),
	}, db.Close
}

type DataAccess struct {
	user     User
	platform Platform
}
