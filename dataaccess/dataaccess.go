package dataaccess

import (
	"Closer/dataaccess/sqliteImpl"
	"fmt"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DataAccess struct {
	User     User
	Platform Platform
}

func NewDataAccess(dbPath string) (*DataAccess, func() error) {
	db, err := gorm.Open("sqlite3", dbPath)
	if err != nil {
		panic(fmt.Sprintf("failed to open DB connection with err %s", err.Error()))
	}

	return &DataAccess{
		Platform: sqliteImpl.NewPlatformDB(db),
		User:     sqliteImpl.NewUserDB(db),
	}, db.Close
}
