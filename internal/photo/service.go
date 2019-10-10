package photo

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"github.com/gin-gonic/gin"
)

type photoItem struct {
	Id int `json:"id"`
	Name string `json:"name" binding:"required"`
	Category string `json:"category" binding:"required"`
}

func (p photoItem) String() string {
	return fmt.Sprintf("[id=%d name=%s, category=%s]", p.Id, p.Name, p.Category)
}

var photoDb = make(map[int]photoItem)

func CreatePhoto(c *gin.Context) {

	var item photoItem

	if c.Bind(&item) == nil {
		id := rand.Intn(100)
		item.Id = id
		photoDb[id] = item
		log.Printf("Added %s into the data store", item)
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, id)
	}
}

func ReadAllPhotos(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, photoDb)
	log.Printf("Returned all %d items", len(photoDb))
}
