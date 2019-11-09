package controllers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const (pingMsg = "closer API v0.1")

type RootController struct {
	log *log.Logger
}

func NewRootController(l *log.Logger) *RootController {
	return &RootController{log: l}
}

func (c *RootController) RegisterRouting(eng *gin.Engine){
	eng.GET("/ping", c.ping)
	eng.GET("/", c.ping)
}

func (c *RootController) ping(ctx *gin.Context){
	ctx.JSON(http.StatusOK, pingMsg)
}
