package knowledge

import "time"

type Article struct {
	Id         int       `json:"id" db:"id"`
	Title      string    `json:"title" binding:"required" db:"title"`
	Content    string    `json:"content" binding:"required" db:"content"`
	DateCreate time.Time `json:"date_create" db:"date_create"`
	DateUpdate time.Time `json:"date_update" db:"date_update"`
	Author     int       `json:"author" db:"author"`
	Views      int       `json:"views" db:"views"`
	IsHide     bool      `json:"hide" db:"is_hide"`
	Likes      int       `json:"likes" db:"likes"`
	Dislikes   int       `json:"dislikes" db:"dislikes"`
	Parent     *int      `json:"parent" db:"parent"`
}

type ArticlesGroup struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}

type ArticlesGroupArticles struct {
	Id        int
	ArticleId int
	GroupId   int
}
