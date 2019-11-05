package likes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

type LikesService struct {
	db *gorm.DB
}

func NewLikesService(db *gorm.DB, createTables string) *LikesService {
	table := &likesItem{}
	if createTables == "true" {
		db.DropTableIfExists(table)
		db.CreateTable(table)
	}
	return &LikesService{db}
}

func (ls *LikesService) AddLikes(c *gin.Context) {

	var item likesItem
	var savedItem likesItem

	if c.Bind(&item) == nil {
		tx := ls.db.Begin()
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()

		err := tx.Error

		// Insert or Update the likes
		if err == nil {
			if foundErr := tx.First(&savedItem, item.Id).Error; foundErr != nil {
				if gorm.IsRecordNotFoundError(foundErr) {
					if err := ls.db.Create(&item).Error; err == nil {
						err = tx.Commit().Error
						savedItem = item
					}
				} else {
					err = foundErr
				}
			} else {
				savedItem.Likes += item.Likes
				if err := tx.Save(&savedItem).Error; err == nil {
					err = tx.Commit().Error
				}
			}
		}

		if err == nil {
			log.Printf("Updated in data store %s", savedItem)
			c.Header("Content-Type", "application/json")
			c.Status(http.StatusOK)
		} else {
			log.Printf("Failed to update likes. %s", err)
			tx.Rollback()
			c.Status(500)
		}
	}
}

func (ps *LikesService) ReadAllLikes(c *gin.Context) {
	var items []likesItem

	if err := ps.db.Find(&items).Error; err != nil {
		log.Printf("Failed to retrieve all likes. %s", err)
		c.Status(500)
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, items)
	log.Printf("Returned all %d items", len(items))
}
