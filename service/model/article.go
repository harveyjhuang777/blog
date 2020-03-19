package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Slug        string `gorm:"unique_index"`
	Title       string
	Description string `gorm:"size:2048"`
	Body        string `gorm:"size:2048"`
	Author      ArticleUser
	AuthorID    uint
	Tags        []Tag     `gorm:"many2many:article_tags;"`
	Comments    []Comment `gorm:"ForeignKey:ArticleID"`
}

type ArticleUser struct {
	gorm.Model
	User      User
	UserID    uint
	Articles  []Article  `gorm:"ForeignKey:AuthorID"`
	Favorites []Favorite `gorm:"ForeignKey:FavoriteByID"`
}

type Favorite struct {
	gorm.Model
	Favorite     Article
	FavoriteID   uint
	FavoriteBy   ArticleUser
	FavoriteByID uint
}

type Tag struct {
	gorm.Model
	Tag      string    `gorm:"unique_index"`
	Articles []Article `gorm:"many2many:article_tags;"`
}

type Comment struct {
	gorm.Model
	Article   Article
	ArticleID uint
	Author    ArticleUser
	AuthorID  uint
	Body      string `gorm:"size:2048"`
}
