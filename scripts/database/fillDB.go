package main

import (
	"Closer/common/platforms"
	"Closer/common/users"
	"fmt"
	"github.com/jinzhu/gorm"
	"os"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	mockPic     = []byte("c2h0b290IHplIHN0YW0gdG1vb25hIHZlIGxvIHN0cmluZyBhdGEgbWVkYW1pZW4gYXMgZHN3dGcgdDRXIDM1OSVUIFky")
	MockUsersDB = []*users.User{
		{
			FirstName:  "Nitzan",
			LastName:   "Uzan",
			Email:      "nitzan@somthing.com",
			Platforms:  []platforms.Platform{
				{
					Type:           platforms.Google,
					Username:       "Nitzanu",
					Nickname:       "nitz",
					ProfilePicture: mockPic,
				},
				{
					Type:           platforms.Facebook,
					Username:       "Nitzanu",
					Nickname:       "nitz",
					ProfilePicture: mockPic,
				},
			},
		},
		{
			FirstName:  "Maayan",
			LastName:   "Layfer",
			Email:      "maayan@somthing.com",
			Platforms:  []platforms.Platform{
				{
					Type:           platforms.Instagram,
					Username:       "Maayan",
					Nickname:       "Maayan",
					ProfilePicture: mockPic,
				},
			},
		},
		{
			FirstName:  "Lychee",
			LastName:   "The Dog",
			Email:      "woofwoof@gmail.com",
			Platforms:  []platforms.Platform{
				{
					Type:           platforms.Facebook,
					Username:       "Lychee",
					Nickname:       "lycha",
					ProfilePicture: mockPic,
				},
				{
					Type:           platforms.Youtube,
					Username:       "Lychee",
					Nickname:       "lych",
					ProfilePicture: mockPic,
				},
				{
					Type:           platforms.Linkedin,
					Username:       "Lychee",
					Nickname:       "lycheeeeeeeee",
					ProfilePicture: mockPic,
				},
			},
		},
		{
			FirstName:  "Chai",
			LastName:   "The Dog",
			Email:      "woofwoof2@gmail.com",
			Platforms:  []platforms.Platform{
				{
					Type:           platforms.Instagram,
					Username:       "chai1",
					Nickname:       "chacha",
					ProfilePicture: mockPic,
				},
			},
		},
	}
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

	for _, u := range MockUsersDB {
		err := db.Create(u)
		if err != nil {
			fmt.Println(err)
		}
	}
}