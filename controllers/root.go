package controllers

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

type RootController struct {
	log *log.Logger
}

func NewRootController(l *log.Logger) RootController {
	return RootController{log: l}
}

func (c RootController) ServeHTTP(w http.ResponseWriter, t *http.Request) {
	_, e := w.Write([]byte("closer API v0.1"))
	if e != nil {
		c.log.Println("ERROR: failed responding to request at rootController")
	}
}
