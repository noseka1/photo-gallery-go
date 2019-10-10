package likes

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

type likesItem struct {
	Id int `json:"id" binding:"required"`
	Likes int `json:"likes" binding:"required"`
}

func (p likesItem) String() string {
	return fmt.Sprintf("[id=%d likes=%d]", p.Id, p.Likes)
}

var likesDb = make(map[int]likesItem)

func AddLikes(c *gin.Context) {

	var json likesItem

	if c.Bind(&json) == nil {
		item, found := likesDb[json.Id]
		if found {
			item.Likes += json.Likes
		} else {
			item = json
		}
		likesDb[item.Id] = item
		log.Printf("Updated in data store %s", item)
		c.Header("Content-Type", "application/json")
		c.Status(http.StatusOK)
	}
}

func ReadAllLikes(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, likesDb)
	log.Printf("Returned all %d items", len(likesDb))
}