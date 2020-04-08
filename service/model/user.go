package model

type User struct {
	ID       uint    `json:"-" gorm:"primary_key"`
	Username string  `json:"username" gorm:"column:username;unique_index"`
	Password string  `json:"password" gorm:"column:password;not null"`
	Email    string  `json:"email" gorm:"column:email;unique_index"`
	Bio      string  `json:"bio" gorm:"column:bio;size:1024"`
	Image    *string `json:"image" gorm:"column:image"`
}

func (User) TableName() string {
	return "user"
}

type UserGetCond struct {
	Username *string `json:"username" gorm:"column:username;unique_index"`
	Email    *string `json:"email" gorm:"column:email;unique_index"`
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
