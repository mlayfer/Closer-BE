package main

import (
	"Closer/common/platforms"
	"Closer/common/users"
	"fmt"
	"github.com/jinzhu/gorm"
	"os"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	dbPath := ".\\CloserDB"
	_ = os.Remove(dbPath)

	db, err := gorm.Open("sqlite3", dbPath)
	if err != nil {
		panic(fmt.Sprintf("failed to open DB connection with err %s", err.Error()))
	}
	defer db.Close()

	db.DropTableIfExists(platforms.Platform{})
	db.DropTableIfExists(users.User{})

	db.AutoMigrate(users.User{}, platforms.Platform{})
}
