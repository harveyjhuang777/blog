package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jwjhuang/blog/service/app/logger"
	"github.com/jwjhuang/blog/service/utils/auth"
	"github.com/jwjhuang/blog/service/utils/errs"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := c.Request
		reqToken := r.Header.Get("Authorization")
		if reqToken == "" {
			logger.Log().Error(errs.TokenMalformed)
			c.JSON(http.StatusUnauthorized, errs.ErrNoAccessToken)
			c.Abort()
			return
		}

		splitToken := strings.Split(reqToken, "Bearer ")
		if len(splitToken) < 2 {
			logger.Log().Error(errs.TokenInvalid)
			c.JSON(http.StatusUnauthorized, errs.TokenInvalid)
			c.Abort()
			return
		}

		tokenString := splitToken[1]
		if tokenString == "" {
			logger.Log().Error(errs.TokenInvalid)
			c.JSON(http.StatusUnauthorized, errs.TokenInvalid)
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				logger.Log().Error(errs.TokenMalformed)
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(auth.SecretKey), nil
		})

		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorExpired != 0 {
					logger.Log().Error(errs.TokenInvalid)
					c.JSON(http.StatusUnauthorized, errs.TokenExpired.Error())
					c.Abort()
					return
				}

				logger.Log().Error(err)
				c.JSON(http.StatusUnauthorized, errs.TokenMalformed)
				c.Abort()
				return
			}
		}

		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("claims", token.Claims)
		}

		c.Next()
	}
}
