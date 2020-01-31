package main

import (
	"Closer/controllers"
	"Closer/dataaccess"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var logger = log.New()

func main() {
	db, dbClose := dataaccess.NewDataAccess(".\\CloserDB")
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

func GetAllRoutables(db *dataaccess.DataAccess) []interface{ RegisterRouting(eng *gin.Engine) } {
	return []interface{ RegisterRouting(eng *gin.Engine) }{
		controllers.NewRootController(logger),
		controllers.NewUsersController(logger, db.User),
		controllers.NewRegistrationController(logger, db),
	}
}
