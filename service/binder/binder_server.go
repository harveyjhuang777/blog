package binder

import (
	"github.com/jwjhuang/blog/service/app/logger"
	"github.com/jwjhuang/blog/service/server"
	"go.uber.org/dig"
)

func provideServe(binder *dig.Container) {
	if err := binder.Provide(server.NewGinServer); err != nil {
		logger.Log().Panic(err.Error())
	}
}
