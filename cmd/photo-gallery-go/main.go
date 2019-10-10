package main

import (
	"os"
	"github.com/gin-gonic/gin"

	"photo-gallery-go/internal/photo"
	"photo-gallery-go/internal/likes"
)

func setupRouter() *gin.Engine {

    // Creates a router without any middleware by default
	router := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// Per route middleware
	router.POST("/photos", photo.CreatePhoto)
	router.GET("/photos", photo.ReadAllPhotos)
	router.POST("/likes", likes.AddLikes)
	router.GET("/likes", likes.ReadAllLikes)

	return router
}

func main() {
	host, exists := os.LookupEnv("GALLERY_HOST")
	if !exists {
		host = "0.0.0.0"
	}

	port, exists := os.LookupEnv("GALLERY_PORT")
	if !exists {
		port = "8080"
	}

	router := setupRouter()
	router.Run(host + ":" + port)
}