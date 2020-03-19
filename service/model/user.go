package model

import "github.com/jinzhu/gorm"

type User struct {
	ID       uint    `json:"-" gorm:"primary_key"`
	Username string  `json:"username" gorm:"column:username;unique_index"`
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

type UserResponse struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Bio      string  `json:"bio"`
	Image    *string `json:"image"`
}

func (s *User) GeUserResponse() *UserResponse {
	return &UserResponse{
		Username: s.Username,
		Email:    s.Email,
		Bio:      s.Bio,
		Image:    s.Image,
	}
}

type ProfileResponse struct {
	ID        uint    `json:"-"`
	Username  string  `json:"username"`
	Bio       string  `json:"bio"`
	Image     *string `json:"image"`
	Following bool    `json:"following"`
}

func (s *User) GeProfileResponse(isFollowing bool) *ProfileResponse {
	return &ProfileResponse{
		ID:        s.ID,
		Username:  s.Username,
		Bio:       s.Bio,
		Image:     s.Image,
		Following: isFollowing,
	}
}
