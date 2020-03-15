package controller

import (
	"github.com/jwjhuang/blog/service/core/user"
	"go.uber.org/dig"
)

type ginControllerSet struct {
	dig.In

	User user.IUserCenter
}

func NewGinController(set ginControllerSet) *GinController {
	return &GinController{
		User:   newUserController(set.User),
		Health: newHealthController(),
	}
}

type GinController struct {
	//AuthToken IAuthTokenController
	User   IUserController
	Health IHealthController
}
