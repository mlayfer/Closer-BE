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
				PlatformData: &platforms.PlatformData{
					Username:   "Nitzanu",
					Nickname:   "nitz",
					ProfilePic: mockPic,
				},
			},
			{
				Type: platforms.Facebook,
				PlatformData: &platforms.PlatformData{
					Username:   "Nitzanu",
					Nickname:   "nitz",
					ProfilePic: mockPic,
				},
			},
		}),
		users.NewUser("Maayan", "Layfer", "maayan@somthing.com", []*platforms.Platform{
			{
				Type: platforms.Instagram,
				PlatformData: &platforms.PlatformData{
					Username:   "Maayan",
					Nickname:   "Maayan",
					ProfilePic: mockPic,
				},
			},
		}),
		users.NewUser("Lychee", "The Dog", "woofwoof@gmail.com", []*platforms.Platform{
			{
				Type: platforms.Facebook,
				PlatformData: &platforms.PlatformData{
					Username:   "Lychee",
					Nickname:   "lycha",
					ProfilePic: mockPic,
				},
			},
			{
				Type: platforms.Youtube,
				PlatformData: &platforms.PlatformData{
					Username:   "Lychee",
					Nickname:   "lych",
					ProfilePic: mockPic,
				},
			},
			{
				Type: platforms.Linkedin,
				PlatformData: &platforms.PlatformData{
					Username:   "Lychee",
					Nickname:   "lycheeeeeeeee",
					ProfilePic: mockPic,
				},
			},
		}),
		users.NewUser("Chai", "The Dog", "woofwoof2@gmail.com", []*platforms.Platform{
			{
				Type: platforms.Instagram,
				PlatformData: &platforms.PlatformData{
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
	eng.DELETE("/users/:identifier", c.deleteUser)
}

// getUsers godoc
// @Summary get Users
// @Description returns all users in DB
// @Tags users
// @Produce json
// @Success 200 {JSON} string AllUsers
// @Router /users [GET]
func (c *UsersController) getUsers(ctx *gin.Context){
	ctx.JSON(http.StatusOK, mockUsersDB)
}

// deleteUser godoc
// @Summary deletes a user
// @Description removes a user from DB
// @Tags users
// @Receive json
// @Produce json
// @Success 200 {JSON} string theUser
// @Router /users/:identifier [DELETE]
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


// getUserByID godoc
// @Summary returns a specific user
// @Tags users
// @Receive json
// @Produce json
// @Success 200 {JSON} string theUser
// @Router /users/:identifier [POST]
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
