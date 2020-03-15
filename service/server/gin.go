package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jwjhuang/blog/service/app/logger"
	"github.com/jwjhuang/blog/service/controller"
	"github.com/jwjhuang/blog/service/middleware"
)

func NewGinServer(server *gin.Engine, controller *controller.GinController) IGinServer {
	return &GinServer{
		server:     server,
		controller: controller,
	}
}

type IGinServer interface {
	Run()
}

type GinServer struct {
	server     *gin.Engine
	controller *controller.GinController
}

func (srv *GinServer) Run() {
	logger.Log().Info("Start HotelServer")
	srv.globMiddleware()
	srv.router()
	srv.server.Run(":8888")
}

func (srv *GinServer) globMiddleware() {
	srv.server.Use(middleware.Cors())
	srv.server.Use(middleware.RequestID())
}

func (srv *GinServer) router() {
	v1 := srv.server.Group("/v1")
	{
		v1.GET("/health", srv.controller.Health.HealthCheck)

		user := v1.Group("/users")
		{
			user.POST("/login", srv.controller.User.Login)
		}

	}
}
