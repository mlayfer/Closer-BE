package main

import (
	"Closer/controllers"
	"Closer/dataaccess"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	_ "github.com/swaggo/gin-swagger/example/basic/docs"
)

var logger = log.New()

func main() {
	db, dbClose := dataaccess.NewDatabaseAccess("root", "root")
	defer dbClose()

	eng := gin.Default()
	for _, v := range GetAllRoutables(db) {
		v.RegisterRouting(eng)
	}

	e := eng.Run() // listen and serve on localhost:8080
	if e != nil {
		logger.Fatal(e)
	}
}

func GetAllRoutables(db *dataaccess.Database) []interface {RegisterRouting(eng *gin.Engine)} {
	return []interface {RegisterRouting(eng *gin.Engine)}{
		controllers.NewRootController(logger),
		controllers.NewUsersController(logger, db),
		controllers.NewRegistrationController(logger, db),
	}
}
