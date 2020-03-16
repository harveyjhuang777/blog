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
	api := srv.server.Group("/api")
	{
		api.GET("/health", srv.controller.Health.HealthCheck)

		users := api.Group("/users")
		{
			users.POST("/login", srv.controller.User.Login)
			users.POST("", srv.controller.User.Register)
		}

		user := api.Group("/user")
		user.Use(middleware.JWTAuth())
		user.Use(middleware.AuthCheckRole())
		{
			user.GET("", srv.controller.User.Get)
			user.PUT("", srv.controller.User.Update)
		}

		profile := api.Group("/profiles")
		profile.Use(middleware.JWTAuth())
		profile.Use(middleware.AuthCheckRole())
		{
			profile.GET("", srv.controller.User.GetProfile)
			profile.POST("/:username/follow", srv.controller.User.Follow)
			profile.DELETE("/:username/follow", srv.controller.User.UnFollow)
		}
	}
}
