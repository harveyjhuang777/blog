package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Slug        string `gorm:"unique_index"`
	Title       string
	Description string `gorm:"size:2048"`
	Body        string `gorm:"size:2048"`
	Tags        []*Tag `gorm:"many2many:article_tags"`
}

func (Article) TableName() string {
	return "article"
}

type Tag struct {
	Tag      string     `gorm:"primary_key"`
	Articles []*Article `gorm:"many2many:article_tags"`
}

func (Tag) TableName() string {
	return "tag"
}

type ArticleGetCond struct {
	Paging
	Tag       *string `form:"tag" json:"tag"`
	Author    *string `form:"author" json:"author"`
	Favorited *string `form:"favorited" json:"favorited"`
}

type ArticleCreateCond struct {
	Title       *string  `json:"title" form:"title"`
	Description *string  `json:"description" form:"description"`
	Body        *string  `json:"body" form:"body"`
	TagList     []string `json:"tagList" form:"tagList"`
}
