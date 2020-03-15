package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jwjhuang/blog/service/app/logger"
)

func newHealthController() IHealthController {
	return &healthController{}
}

type IHealthController interface {
	HealthCheck(c *gin.Context)
}

type healthController struct {
}

func (rc *healthController) HealthCheck(c *gin.Context) {
	logger.Log().Infof("health check on %v", time.Now().String())
	c.JSON(http.StatusOK, nil)
}
