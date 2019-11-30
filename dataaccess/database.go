package dataaccess

import (
	"Closer/common/platforms"
	"Closer/common/users"
	"fmt"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

type database struct {
	DB *gorm.DB
}

func NewDatabaseAccess(user, pass string) (dbConn *database, closer func()) {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@/CloserDB?charset=utf8&parseTime=True&loc=Local", user, pass))
	if err != nil {
		panic("failed to open DB connection")
	}

	if !db.HasTable("users_platforms") {
		db.Table("users_platforms").CreateTable(new(platforms.Platform))
	} else {
		db.AutoMigrate(new(platforms.Platform))
	}
	if !db.HasTable(new(users.User)) {
		db.Table("users").CreateTable(new(users.User))
	} else {
		db.AutoMigrate(new(users.User))
	}

	return &database{DB: db}, func() { db.Close() }
}
