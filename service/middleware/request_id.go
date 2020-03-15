package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Request-Id", uuid.NewV4().String())
		c.Next()
	}
}
