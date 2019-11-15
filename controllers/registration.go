package controllers

import (
	"Closer/common/users"
	"Closer/platforms"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type RegistrationController struct {
	log *log.Logger
}

func NewRegistrationController(l *log.Logger) *RegistrationController {
	return &RegistrationController{log: l}
}

func (c *RegistrationController) RegisterRouting(eng *gin.Engine){
	eng.POST("/register", c.registerUser)
}

// registerUser godoc
// @Summary adds a user
// @Description add a user to DB
// @Tags users
// @Receive json
// @Produce json
// @Success 200 {JSON} string newUser
// @Router /users/ [POST]
func (c *RegistrationController) registerUser(ctx *gin.Context){
	var ru registerUserRequest
	if err := ctx.BindJSON(&ru); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if (ru.Platforms == nil || len(ru.Platforms) == 0){
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("cannot register without a platform"))
		return
	}

 	u := users.NewUser(ru.FirstName, ru.LastName, ru.Email, ru.Platforms)

	mockUsersDB = append(mockUsersDB, u)
	ctx.JSON(http.StatusOK, gin.H{"identifier": u.Identifier})
}

type registerUserRequest struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Platforms []*platforms.Platform `json:"platforms" binding:"required"`
}