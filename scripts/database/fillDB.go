package main

import (
	"Closer/common/platforms"
	"Closer/common/users"
	"Closer/dataaccess"
	"fmt"
)

var (
	mockPic     = []byte("c2h0b290IHplIHN0YW0gdG1vb25hIHZlIGxvIHN0cmluZyBhdGEgbWVkYW1pZW4gYXMgZHN3dGcgdDRXIDM1OSVUIFky")
	MockUsersDB = []*users.User{
		users.NewUser("Nitzan", "Uzan", "nitzan@somthing.com", []*platforms.Platform{
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
		}),
		users.NewUser("Maayan", "Layfer", "maayan@somthing.com", []*platforms.Platform{
			{
				Type:           platforms.Instagram,
				Username:       "Maayan",
				Nickname:       "Maayan",
				ProfilePicture: mockPic,
			},
		}),
		users.NewUser("Lychee", "The Dog", "woofwoof@gmail.com", []*platforms.Platform{
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
		}),
		users.NewUser("Chai", "The Dog", "woofwoof2@gmail.com", []*platforms.Platform{
			{
				Type:           platforms.Instagram,
				Username:       "chai1",
				Nickname:       "chacha",
				ProfilePicture: mockPic,
			},
		}),
	}
)

func main() {
	closerDB, closer := dataaccess.NewDatabaseAccess("root", "root")
	defer closer()

	for _, u := range MockUsersDB {
		err := closerDB.InsertUser(u)
		if err != nil {
			fmt.Println(err)
		}
	}
}