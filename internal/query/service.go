package query

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

type QueryService struct {
	db *gorm.DB
}

func NewQueryService(db *gorm.DB) *QueryService {
	return &QueryService{db}
}

func (qs *QueryService) ReadCategoryOrderedByLikes(c *gin.Context) {

	category := c.Query("category")

	var items []queryItem

	rows, err := qs.db.Table("photo_items").
		Select("photo_items.id, photo_items.name, photo_items.category, likes_items.likes").
		Joins("left join likes_items on photo_items.id = likes_items.id").
		Where("photo_items.category = ?", category).
		Order("likes_items.likes").
		Rows()
	if err != nil {
		log.Printf("Failed to retrieve photos in the category %s. %s", category, err)
		c.Status(500)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var item queryItem
		rows.Scan(&item.Id, &item.Name, &item.Category, &item.Likes)
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Failed to retrieve rows. %s", err)
		c.Status(500)
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, items)
	log.Printf("Returned all %d items", len(items))
}
