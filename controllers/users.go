package controllers

import (
	"Closer/common/users"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type usersDB interface {
	GetAllUsers() ([]*users.User, error)
	DeleteUser(id uuid.UUID) error
	GetUserByID(id uuid.UUID) (*users.User, error)
	InsertUser(user *users.User) error
}

type UsersController struct {
	log *log.Logger
	db usersDB
}

func NewUsersController(l *log.Logger, database usersDB) *UsersController {
	return &UsersController{log: l, db: database}
}

func (c *UsersController) RegisterRouting(eng *gin.Engine) {
	eng.GET("/users", c.getUsers)
	eng.GET("/users/:identifier", c.getUserByID)
	eng.DELETE("/users/:identifier", c.deleteUser)
}

func (c *UsersController) getUsers(ctx *gin.Context) {
	allUsers, err := c.db.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, allUsers)
}

func (c *UsersController) deleteUser(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("identifier"))
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err = c.db.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.Status(http.StatusNoContent)
	return
}

func (c *UsersController) getUserByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("identifier"))
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	usr, err := c.db.GetUserByID(id)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, usr)
	return
}
