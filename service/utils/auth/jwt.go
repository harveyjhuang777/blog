package auth

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jwjhuang/blog/service/model"
)

var (
	SecretKey string = "hotel"
)

func GenerateJWTToken(user *model.User, secretKey string) (token string, err error) {
	nowTime := time.Now()
	jwtToken := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["id"] = strconv.Itoa(user.ID)
	claims["account"] = user.Account
	claims["role"] = user.Role
	claims["exp"] = nowTime.Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = nowTime.Unix()
	jwtToken.Claims = claims

	token, err = jwtToken.SignedString([]byte(secretKey))
	if err != nil {
		return
	}

	return
}
