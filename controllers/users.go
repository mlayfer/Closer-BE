package controllers

import (
	"Closer/dataaccess"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type UsersController struct {
	log *log.Logger
	db  dataaccess.User
}

func NewUsersController(l *log.Logger, database dataaccess.User) *UsersController {
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
	id, err := uuid.FromString(ctx.Param("identifier"))
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
	id, err := uuid.FromString(ctx.Param("identifier"))
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	usr, err := c.db.GetByIdentifier(id)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, usr)
	return
}
