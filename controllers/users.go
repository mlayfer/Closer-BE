package controllers

import (
	"Closer/common/platforms"
	"Closer/common/users"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"net/http"
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

type UsersController struct {
	log *log.Logger
}

func NewUsersController(l *log.Logger) *UsersController {
	return &UsersController{log: l}
}

func (c *UsersController) RegisterRouting(eng *gin.Engine) {
	eng.GET("/users", c.getUsers)
	eng.GET("/users/:identifier", c.getUserByID)
	eng.DELETE("/users/:identifier", c.deleteUser)
}

func (c *UsersController) getUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, MockUsersDB)
}

func (c *UsersController) deleteUser(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("identifier"))
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	for i, u := range MockUsersDB {
		if u.Identifier == id {
			MockUsersDB = append(MockUsersDB[:i], MockUsersDB[i+1:]...)
			ctx.JSON(http.StatusOK, u)
			return
		}
	}

	ctx.Status(http.StatusNotFound)
	return
}

func (c *UsersController) getUserByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("identifier"))
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	for _, u := range MockUsersDB {
		if u.Identifier == id {
			ctx.JSON(http.StatusOK, u)
			return
		}
	}

	ctx.Status(http.StatusNotFound)
	return
}
