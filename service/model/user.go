package model

import "github.com/jinzhu/gorm"

type User struct {
	ID       uint    `json:"-" gorm:"primary_key"`
	Username string  `json:"username" gorm:"column:username"`
	Password string  `json:"password" gorm:"column:password;not null"`
	Email    string  `json:"email" gorm:"column:email;unique_index"`
	Bio      string  `json:"bio" gorm:"column:bio;size:1024"`
	Image    *string `json:"image" gorm:"column:image"`
}

type Follow struct {
	gorm.Model
	Following    User `json:"following"`
	FollowingID  uint `json:"followingId"`
	FollowedBy   User `json:"followedBy"`
	FollowedByID uint `json:"followedById"`
}
