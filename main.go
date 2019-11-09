package main

import (
	"Closer/controllers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	//"github.com/gin-gonic/gin"
)

func main() {
	logger := log.New()


	//eng := gin.Default()
	//defineRouting(eng)
	//log.Fatal(eng.Run()) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/", controllers.NewRootController(logger))
	router.Handle("/users/", controllers.NewUsersController(logger))
	log.Fatal(http.ListenAndServe(":8080", router))
}

//func defineRouting(eng *gin.Engine){
//	eng.GET()
//	eng.GET("/ping", func(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"message": "pong",
//		})
//	})
//}