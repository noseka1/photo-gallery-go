package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"

	"photo-gallery-go/internal/likes"
	"photo-gallery-go/internal/photo"
)

func setupRouter(ps *photo.PhotoService, ls *likes.LikesService) *gin.Engine {

	// Creates a router without any middleware by default
	router := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// Per route middleware
	router.POST("/photos", ps.CreatePhoto)
	router.GET("/photos", ps.ReadAllPhotos)
	router.POST("/likes", ls.AddLikes)
	router.GET("/likes", ls.ReadAllLikes)

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
	dbHost, exists := os.LookupEnv("GALLERY_DB_HOST")
	if !exists {
		dbHost = "localhost"
	}
	dbPort, exists := os.LookupEnv("GALLERY_DB_PORT")
	if !exists {
		dbPort = "5432"
	}
	dbSsl, exists := os.LookupEnv("GALLERY_DB_SSL")
	if !exists {
		dbSsl = "disable"
	}
	dbName, exists := os.LookupEnv("GALLERY_DB_NAME")
	if !exists {
		dbName = "gallery"
	}
	dbUser, exists := os.LookupEnv("GALLERY_DB_USER")
	if !exists {
		dbUser = "gallery"
	}
	dbPassword, exists := os.LookupEnv("GALLERY_DB_PASSWORD")
	if !exists {
		dbPassword = "password"
	}

	dbConnection := fmt.Sprintf("host=%s port=%s sslmode=%s dbname=%s user=%s", dbHost, dbPort, dbSsl, dbName, dbUser)
	dbConnectionFull := fmt.Sprintf("%s password=%s", dbConnection, dbPassword)

	db, err := gorm.Open("postgres", dbConnectionFull)

	if err == nil {
		defer db.Close()
		log.Printf("Connected to the database %s", dbConnection)

		// Enable Logger, show detailed log
		db.LogMode(true)
		db.SetLogger(log.New(os.Stdout, "\r\n", 0))

		ps := photo.NewPhotoService(db)
		ls := likes.NewLikesService(db)
		router := setupRouter(ps, ls)
		router.Run(host + ":" + port)
	} else {
		log.Printf("Cannot connect to the database %s", dbConnection)
		log.Fatal(err)
	}
}
