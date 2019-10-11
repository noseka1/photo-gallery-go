package query

type queryItem struct {
	Id    int `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Category string `json:"category" binding:"required"`
	Likes int `json:"likes" binding:"required"`
}
