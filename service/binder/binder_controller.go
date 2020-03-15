package binder

import (
	"github.com/jwjhuang/blog/service/app/logger"
	"github.com/jwjhuang/blog/service/controller"
	"go.uber.org/dig"
)

func provideController(binder *dig.Container) {
	// Controller
	if err := binder.Provide(controller.NewGinController); err != nil {
		logger.Log().Panic(err.Error())
	}
}
