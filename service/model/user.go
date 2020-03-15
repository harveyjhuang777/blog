package model

type User struct {
	ID       int    `json:"id"`
	Account  string `firestore:"account" form:"account"  json:"account"`
	Name     string `firestore:"name" form:"name"  json:"name"`
	Password string `firestore:"password" form:"password"  json:"password"`
	Role     string `firestore:"role" form:"role"  json:"role"`
}
