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


	db.AutoMigrate(users.User{})
	db.AutoMigrate(platforms.Platform{})
	//_ = db.Exec("CREATE TABLE Users (	UserIdentifier BLOB   PRIMARY KEY	UNIQUE,		FirstName      STRING NOT NULL,		LastName       STRING NOT NULL,		Email          STRING NOT NULL	);")
	//_ = db.Exec("CREATE TABLE Platforms ( UserIdentifier BLOB (16) REFERENCES Users (UserID) NOT NULL,Identifier     BLOB      PRIMARY KEY	UNIQUE,		Type           INT       NOT NULL,		UserName       STRING,		Password       STRING,		Nickname       STRING,		ProfilePic     BLOB	);	")
}
