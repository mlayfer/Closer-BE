package controllers

import (
	"Closer/common/users"
	"Closer/platforms"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func NewUsersController(l *log.Logger) *UsersController {
	return &UsersController{log: l}
}

func (c *UsersController) RegisterRouting(eng *gin.Engine){
	eng.GET("/users", c.getUsers)
	eng.GET("/users/:identifier", c.getUserByID)
	eng.POST("/users", c.addUser)
	eng.DELETE("/users/:identifier", c.deleteUser)
}

func (c *UsersController) getUsers(ctx *gin.Context){
	ctx.JSON(http.StatusOK, mockUsersDB)
}

func (c *UsersController) deleteUser(ctx *gin.Context){
	id, err := uuid.Parse(ctx.Param("identifier"))
	if err != nil{
		ctx.Status(http.StatusBadRequest)
		return
	}

	for i, u := range mockUsersDB {
		if u.Identifier ==  id{
			mockUsersDB = append(mockUsersDB[:i], mockUsersDB[i+1:]...)
			ctx.JSON(http.StatusOK, u)
			return
		}
	}

	ctx.Status(http.StatusNotFound)
	return
}

func (c *UsersController) addUser(ctx *gin.Context){
	// Parse JSON
	var json struct {
		Value users.User `json:"value" binding:"required"`
	}

	if ctx.Bind(&json) == nil {
		mockUsersDB = append(mockUsersDB, &json.Value)
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}

func (c *UsersController) getUserByID(ctx *gin.Context){
	id, err := uuid.Parse(ctx.Param("identifier"))
	if err != nil{
		ctx.Status(http.StatusBadRequest)
		return
	}

	for _, u := range mockUsersDB {
		if u.Identifier ==  id{
			ctx.JSON(http.StatusOK, u)
			return
		}
	}

	ctx.Status(http.StatusNotFound)
	return
}
