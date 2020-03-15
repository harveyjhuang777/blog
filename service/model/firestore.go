package model

import "time"

type FireStoreBase struct {
	ID        string     `firestore:"id" form:"id" json:"id"`
	CreatedAt time.Time  `firestore:"createdAt" form:"createdAt" json:"-"`
	UpdatedAt time.Time  `firestore:"updatedAt" form:"updatedAt" json:"-"`
	DeletedAt *time.Time `firestore:"deletedAt" form:"deletedAt" json:"-"`
}
