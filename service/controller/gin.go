package controller

import (
	"github.com/jwjhuang/blog/service/core/article"
	"github.com/jwjhuang/blog/service/core/user"
	"go.uber.org/dig"
)

type ginControllerSet struct {
	dig.In

	User    user.IUserCenter
	Article article.IArticleCenter
}

func NewGinController(set ginControllerSet) *GinController {
	return &GinController{
		User:    newUserController(set.User),
		Health:  newHealthController(),
		Article: newArticleController(set.Article),
	}
}

type GinController struct {
	User    IUserController
	Health  IHealthController
	Article IArticleController
}
