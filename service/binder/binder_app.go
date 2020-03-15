package binder

import (
	"github.com/jwjhuang/blog/service/app/framework"
	"github.com/jwjhuang/blog/service/app/logger"
	"github.com/jwjhuang/blog/service/app/storage"
	"github.com/jwjhuang/blog/service/utils/auth"
	"go.uber.org/dig"
)

func provideApp(binder *dig.Container) {
	if err := binder.Provide(storage.NewFirestore); err != nil {
		logger.Log().Panic(err.Error())
	}
	if err := binder.Provide(framework.NewGin); err != nil {
		logger.Log().Panic(err.Error())
	}
	if err := binder.Provide(storage.NewGORM); err != nil {
		logger.Log().Panic(err.Error())
	}
	if err := binder.Invoke(auth.NewAccess); err != nil {
		logger.Log().Panic(err.Error())
	}
}
