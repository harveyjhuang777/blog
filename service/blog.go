package service

import (
	"github.com/jwjhuang/blog/service/app/logger"
	"github.com/jwjhuang/blog/service/binder"
	"github.com/jwjhuang/blog/service/server"
	"go.uber.org/dig"
)

type Serve int

const (
	ServeBlog Serve = iota
)

var (
	serverStrategy = map[Serve]IServe{
		ServeBlog: newBlogServer(),
	}
)

type IServe interface {
	Run()
}

func Run(srv Serve) {
	logger.Start()
	obj, ok := serverStrategy[srv]
	if !ok {
		logger.Log().Panic("server strategy not found")
	}
	obj.Run()

	select {}
}

func newBlogServer() IServe {
	return &BlogServer{}
}

type BlogServer struct {
}

func (hs BlogServer) Run() {
	binder := binder.New()
	if err := binder.Invoke(hs.gen); err != nil {
		logger.Log().Panic(err.Error())
	}
}

func (hs BlogServer) gen(set serverSet) {
	go set.BlogServer.Run()
}

type serverSet struct {
	dig.In

	BlogServer server.IGinServer
}
