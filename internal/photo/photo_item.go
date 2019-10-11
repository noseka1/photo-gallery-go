package photo

import (
	"fmt"
)

type photoItem struct {
	Id       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	Category string `json:"category" binding:"required"`
}

func (p photoItem) String() string {
	return fmt.Sprintf("[id=%d name=%s, category=%s]", p.Id, p.Name, p.Category)
}
