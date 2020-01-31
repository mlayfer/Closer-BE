package controllers

import (
	"Closer/common/platforms"
	"Closer/common/users"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type RegistrationController struct {
	log *log.Logger
	db  usersDB
}

func NewRegistrationController(l *log.Logger, database usersDB) *RegistrationController {
	return &RegistrationController{log: l, db: database}
}

func (c *RegistrationController) RegisterRouting(eng *gin.Engine) {
	eng.POST("/register", c.registerUser)
}

func (c *RegistrationController) registerUser(ctx *gin.Context) {
	var ru registerUserRequest
	if err := ctx.BindJSON(&ru); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if ru.Platforms == nil || len(ru.Platforms) == 0 {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("cannot register without a platform"))
		return
	}

	u := &users.User{
		FirstName: ru.FirstName,
		LastName:  ru.LastName,
		Email:     ru.Email,
		Platforms: ru.Platforms,
	}
	if err := c.db.InsertUser(u); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"identifier": u.Identifier})
}

type registerUserRequest struct {
	FirstName string                `json:"firstname" binding:"required"`
	LastName  string                `json:"lastname" binding:"required"`
	Email     string                `json:"email" binding:"required"`
	Platforms []*platforms.Platform `json:"platforms" binding:"required"`
}
