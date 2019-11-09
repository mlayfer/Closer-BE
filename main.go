package main

import (
	"Closer/controllers"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	logger := log.New()

	eng := gin.Default()
	for _, v := range GetAllRoutables(logger) {
		v.RegisterRouting(eng)
	}
	e := eng.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if e != nil {
		logger.Fatal(e)
	} else {
		return
	}
}

func GetAllRoutables(l *log.Logger) []interface {
	RegisterRouting(eng *gin.Engine)
}{
	return []interface {
		RegisterRouting(eng *gin.Engine)
	}{
		controllers.NewRootController(l),
		controllers.NewUsersController(l),
	}
}