package photo

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

type PhotoService struct {
	db *gorm.DB
}

func NewPhotoService(db *gorm.DB, createTables string) *PhotoService {
	table := &photoItem{}
	if (createTables == "true") {
		db.DropTableIfExists(table)
		db.CreateTable(table)
	}
	return &PhotoService{db}
}

func (ps *PhotoService) CreatePhoto(c *gin.Context) {

	var item photoItem

	if c.Bind(&item) == nil {
		if err := ps.db.Create(&item).Error; err != nil {
			log.Printf("Failed to create photo. %s", err)
			c.Status(500)
			return
		}
		log.Printf("Added %s into the data store", item)
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, item.Id)
	}
}

func (ps *PhotoService) ReadAllPhotos(c *gin.Context) {

	var items []photoItem

	if err := ps.db.Find(&items).Error; err != nil {
		log.Printf("Failed to retrieve all photos. %s", err)
		c.Status(500)
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, items)
	log.Printf("Returned all %d items", len(items))
}
