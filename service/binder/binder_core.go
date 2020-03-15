package binder

import (
	"github.com/jwjhuang/blog/service/app/logger"
	"github.com/jwjhuang/blog/service/core/user"
	"go.uber.org/dig"
)

func provideCore(binder *dig.Container) {
	if err := binder.Provide(user.NewUserCenter); err != nil {
		logger.Log().Panic(err.Error())
	}
}
