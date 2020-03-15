package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jwjhuang/blog/service/utils/auth"
	"github.com/jwjhuang/blog/service/utils/errs"
)

func AuthCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user interface{}

		claims, ok := c.Get("claims")
		if ok {
			user = claims.(jwt.MapClaims)["account"]
		}

		e := auth.GetAccessInstance()
		hasRight, err := e.Enforce(user.(string), c.Request.URL.Path, c.Request.Method)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}
		if hasRight {
			c.Next()
		} else {
			c.JSON(http.StatusForbidden, errs.ErrNoAccessRight.Error())
			c.Abort()
			return
		}
	}
}
