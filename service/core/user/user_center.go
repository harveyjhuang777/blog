package user

import (
	"sync"

	"cloud.google.com/go/firestore"
	"go.uber.org/dig"
)

var (
	packet userCenterSet
)

var (
	once sync.Once
	ptr  *userCenter
)

func NewUserCenter(set userCenterSet) userCenter {
	once.Do(func() {
		packet = set

		newDAO()

		ptr = &userCenter{
			User: newUser(),
		}
	})

	return *ptr
}

type userCenterSet struct {
	dig.In

	DB *firestore.Client
}

type userCenter struct {
	dig.Out

	User IUserCenter
}
