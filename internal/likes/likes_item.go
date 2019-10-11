package likes

import (
	"fmt"
)

type likesItem struct {
	Id    int `json:"id" binding:"required"`
	Likes int `json:"likes" binding:"required"`
}

func (p likesItem) String() string {
	return fmt.Sprintf("[id=%d likes=%d]", p.Id, p.Likes)
}