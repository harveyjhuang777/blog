package user

import (
	"sync"

	"github.com/jinzhu/gorm"
	"go.uber.org/dig"
)

var (
	packet userCenterSet
)

var (
	once sync.Once
	ptr  *userCenter
)

//NewUserCenter generate user instance for external usage
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

	DB *gorm.DB
}

type userCenter struct {
	dig.Out

	User IUserCenter
}
