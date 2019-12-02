package dataaccess

import (
	"Closer/common/platforms"
	"Closer/common/users"
	"fmt"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	db *gorm.DB
}

func NewDatabaseAccess(user, pass string) (dbConn *Database, closer func()) {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@/CloserDB?charset=utf8&parseTime=True&loc=Local", user, pass))
	if err != nil {
		panic("failed to open DB connection")
	}

	return &Database{db: db}, func() { db.Close() }
}

func (d *Database) InitDB(){
	if !d.db.HasTable("platforms") {
		d.db.Table("platforms").CreateTable(new(platforms.Platform))
	} else {
		d.db.AutoMigrate(new(platforms.Platform))
	}
	if !d.db.HasTable(new(users.User)) {
		d.db.Table("users").CreateTable(new(users.User))
	} else {
		d.db.AutoMigrate(new(users.User))
	}
}