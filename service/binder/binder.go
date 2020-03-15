package binder

import (
	"github.com/jwjhuang/blog/service/app/logger"
	"go.uber.org/dig"
)

var (
	binder *dig.Container
)

func New() *dig.Container {
	logger.Log().Info("Init DenpendcieS")
	binder = dig.New()

	provideApp(binder)
	provideController(binder)
	provideCore(binder)
	provideServe(binder)

	return binder
}
