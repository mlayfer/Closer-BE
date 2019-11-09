package controllers

import (
	"Closer/common/users"
	"Closer/platforms"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var (
	mockPic     = []byte("c2h0b290IHplIHN0YW0gdG1vb25hIHZlIGxvIHN0cmluZyBhdGEgbWVkYW1pZW4gYXMgZHN3dGcgdDRXIDM1OSVUIFky")
	mockUsersDB = []*users.User{
		users.NewUser("Nitzan", "Uzan", "nitzan@somthing.com", []*platforms.Platform{
			{
				Type: platforms.Google,
				PlatformsData: &platforms.PlatformData{
					Username:   "Nitzanu",
					Nickname:   "nitz",
					ProfilePic: mockPic,
				},
			},
			{
				Type: platforms.Facebook,
				PlatformsData: &platforms.PlatformData{
					Username:   "Nitzanu",
					Nickname:   "nitz",
					ProfilePic: mockPic,
				},
			},
		}),
		users.NewUser("Maayan", "Layfer", "maayan@somthing.com", []*platforms.Platform{
			{
				Type: platforms.Instagram,
				PlatformsData: &platforms.PlatformData{
					Username:   "Maayan",
					Nickname:   "Maayan",
					ProfilePic: mockPic,
				},
			},
		}),
		users.NewUser("Lychee", "The Dog", "woofwoof@gmail.com", []*platforms.Platform{
			{
				Type: platforms.Facebook,
				PlatformsData: &platforms.PlatformData{
					Username:   "Lychee",
					Nickname:   "lycha",
					ProfilePic: mockPic,
				},
			},
			{
				Type: platforms.Youtube,
				PlatformsData: &platforms.PlatformData{
					Username:   "Lychee",
					Nickname:   "lych",
					ProfilePic: mockPic,
				},
			},
			{
				Type: platforms.Linkedin,
				PlatformsData: &platforms.PlatformData{
					Username:   "Lychee",
					Nickname:   "lycheeeeeeeee",
					ProfilePic: mockPic,
				},
			},
		}),
		users.NewUser("Chai", "The Dog", "woofwoof2@gmail.com", []*platforms.Platform{
			{
				Type: platforms.Instagram,
				PlatformsData: &platforms.PlatformData{
					Username:   "chai1",
					Nickname:   "chacha",
					ProfilePic: mockPic,
				},
			},
		}),
	}
)

type UsersController struct {
	log *log.Logger
}

func NewUsersController(l *log.Logger) UsersController {
	return UsersController{log: l}
}

func (c UsersController) ServeHTTP(w http.ResponseWriter, t *http.Request) {
	switch t.Method {
	case "GET":
		j, err := json.Marshal(mockUsersDB)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			break
		}
		w.Write(j)
		break
	}
}

func (c *UsersController) GetUsers() ([]*users.User, error){
	return mockUsersDB, nil
}