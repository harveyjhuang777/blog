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
		{
			user.GET("", srv.controller.User.Get)
			user.PUT("", srv.controller.User.Update)
		}

		profile := api.Group("/profile")
		{
			profile.GET("", srv.controller.User.Profile)
		}

		articles := api.Group("/articles")
		articles.Use(middleware.JWTAuth())
		{
			articles.GET("", srv.controller.Article.List)
			articles.GET("/:slug", srv.controller.Article.Get)
			articles.POST("", srv.controller.Article.Create)
			articles.PUT("/:slug", srv.controller.Article.Update)
			articles.DELETE("/:slug", srv.controller.Article.Delete)
		}

		tags := api.Group("/tags")
		{
			tags.GET("", srv.controller.Article.TagList)
		}
	}
}
