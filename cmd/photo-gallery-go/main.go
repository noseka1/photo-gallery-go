package main

import (
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
	router := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")
}