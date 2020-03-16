package model

import "github.com/jinzhu/gorm"

type User struct {
	ID       uint    `gorm:"primary_key"`
	Username string  `gorm:"column:username"`
	Password string  `gorm:"column:password;not null"`
	Email    string  `gorm:"column:email;unique_index"`
	Bio      string  `gorm:"column:bio;size:1024"`
	Image    *string `gorm:"column:image"`
}

type Follow struct {
	gorm.Model
	Following    User
	FollowingID  uint
	FollowedBy   User
	FollowedByID uint
}
