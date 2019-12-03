package main

import (
	routers "github.com/mazeForGit/WordlistExtractor/routers"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/static"
	log "github.com/sirupsen/logrus"
	"os"
)

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "5000"
	}
	return ":" + port
}

func main() {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	router := gin.Default()
	router.RedirectTrailingSlash = false

	router.LoadHTMLGlob("public/*.html")
	router.Use(static.Serve("/", static.LocalFile("./public", false)))
	router.GET("/", routers.Index)
	router.NoRoute(routers.NotFoundError)
	router.GET("/500", routers.InternalServerError)
	//router.GET("/health", routers.HealthGET)

	router.GET("/words", routers.WordsGET)
	router.POST("/words", routers.WordsPOST)
	router.GET("/words/:id", routers.WordsByIdGET)
	router.DELETE("/words/:id", routers.WordsByIdDELETE)
	router.GET("/tests", routers.TestsGET)
	router.GET("/wordlist", routers.WordListGET)
	router.PUT("/wordlist", routers.WordListPUT)
	router.DELETE("/wordlist", routers.WordListDELETE)

	log.Info("Starting gowebapp on port " + port())

	router.Run(port())
}