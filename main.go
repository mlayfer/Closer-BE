package main

import (
	"Closer/controllers"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "github.com/swaggo/gin-swagger/example/basic/docs"
)

// @title Swagger Closer API
// @version 0.1
// @description This is a server for closer application
// @host localhost:8080
// @BasePath /
func main() {
	eng := gin.Default()

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	eng.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))


	for _, v := range GetAllRoutables() {
		v.RegisterRouting(eng)
	}

	e := eng.Run() // listen and serve on localhost:8080
	if e != nil {
		logger.Fatal(e)
	} else {
		return
	}
}

func GetAllRoutables() []interface {
	RegisterRouting(eng *gin.Engine)
}{
	return []interface {
		RegisterRouting(eng *gin.Engine)
	}{
		controllers.NewRootController(logger),
		controllers.NewUsersController(logger),
		controllers.NewRegistrationController(logger),
	}
}

var logger = log.New()