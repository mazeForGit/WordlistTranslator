package main

import (
	"os"
	
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/static"
	log "github.com/sirupsen/logrus"
	
	routers "github.com/mazeForGit/WordlistTranslator/routers"
	data "github.com/mazeForGit/WordlistTranslator/model"
)

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func main() {
	
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	router := gin.Default()
	router.RedirectTrailingSlash = false

	//pprof.Register(router)

	router.LoadHTMLGlob("public/*.html")
	router.Use(static.Serve("/", static.LocalFile("./public", false)))
	router.GET("/", routers.Index)
	router.GET("/index", routers.Index)
	router.NoRoute(routers.NotFoundError)
	router.GET("/500", routers.InternalServerError)
	router.GET("/health", routers.HealthGET)

	router.GET("/config", routers.ConfigGET)
	router.POST("/config", routers.ConfigPOST)
	router.PUT("/config", routers.ConfigPUT)

	log.Info("Starting background process")
	go model.ExecuteLongRunningTaskOnRequest()
	
	log.Info("Starting gowebapp on port " + port())
	router.Run(port())
	
}