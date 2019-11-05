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
	"photo-gallery-go/internal/query"
)

type config struct {
	host           string
	port           string
	dbHost         string
	dbPort         string
	dbSsl          string
	dbName         string
	dbUser         string
	dbPassword     string
	dbCreateTables bool
}

func readConfiguration() *config {

	conf := &config{
		host:           "0.0.0.0",
		port:           "8080",
		dbHost:         "localhost",
		dbPort:         "5432",
		dbSsl:          "disable",
		dbName:         "gallery",
		dbUser:         "gallery",
        dbPassword:     "password",
	    dbCreateTables: false,
	}

	var tmp string
	var exists bool

	tmp, exists = os.LookupEnv("GALLERY_HOST")
	if exists {
		conf.host = tmp
	}
	tmp, exists = os.LookupEnv("GALLERY_PORT")
	if exists {
		conf.port = tmp
	}
	tmp, exists = os.LookupEnv("GALLERY_DB_HOST")
	if exists {
		conf.dbHost = tmp
	}
	tmp, exists = os.LookupEnv("GALLERY_DB_PORT")
	if exists {
		conf.dbPort = tmp
	}
	tmp, exists = os.LookupEnv("GALLERY_DB_SSL")
	if exists {
		conf.dbSsl = tmp
	}
	tmp, exists = os.LookupEnv("GALLERY_DB_NAME")
	if exists {
		conf.dbName = tmp
	}
	tmp, exists = os.LookupEnv("GALLERY_DB_USER")
	if exists {
		conf.dbUser = tmp
	}
	tmp, exists = os.LookupEnv("GALLERY_DB_PASSWORD")
	if exists {
		conf.dbPassword = tmp
	}
	tmp, exists = os.LookupEnv("GALLERY_DB_CREATE_TABLES")
	if exists && tmp == "true" {
		conf.dbCreateTables = true
	}
	return conf
}

func setupRouter(ps *photo.PhotoService, ls *likes.LikesService, qs *query.QueryService) *gin.Engine {

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
	router.GET("/query", qs.ReadCategoryOrderedByLikes)

	return router
}

func main() {

	conf := readConfiguration()

	dbConnection := fmt.Sprintf("host=%s port=%s sslmode=%s dbname=%s user=%s", conf.dbHost, conf.dbPort, conf.dbSsl, conf.dbName, conf.dbUser)
	dbConnectionFull := fmt.Sprintf("%s password=%s", dbConnection, conf.dbPassword)

	// Open database connection
	db, err := gorm.Open("postgres", dbConnectionFull)
	if err != nil {
		log.Printf("Cannot connect to the database %s", dbConnection)
		log.Fatal(err)
		os.Exit(1)
	}
	defer db.Close()
	log.Printf("Connected to the database %s", dbConnection)

	// Enable Logger, show detailed log
	db.LogMode(true)
	db.SetLogger(log.New(os.Stdout, "\r\n", 0))

	// Create services
	ps := photo.NewPhotoService(db)
	ls := likes.NewLikesService(db)
	qs := query.NewQueryService(db)

	// Drop and create database tables
	if conf.dbCreateTables {
		ps.CreateTables()
		ls.CreateTables()
	}

	// Connect services to the API and start the router
	router := setupRouter(ps, ls, qs)
	router.Run(conf.host + ":" + conf.port)
}
