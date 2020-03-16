package user

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jwjhuang/blog/service/app/logger"
	"github.com/jwjhuang/blog/service/model"
	"github.com/jwjhuang/blog/service/utils/auth"
	"golang.org/x/crypto/bcrypt"
)

func newUser() IUserCenter {
	return &userUseCase{}
}

type IUserCenter interface {
	Login(c *gin.Context) (*model.Token, error)
	Take(c *gin.Context) (*model.User, error)
}

type userUseCase struct {
}

func (uc *userUseCase) Login(c *gin.Context) (*model.Token, error) {

	user := &model.User{}
	err := c.ShouldBindBodyWith(user, binding.JSON)
	if err != nil {
		logger.Log().Error(err)
		return nil, err
	}

	dbUser, err := uc.Take(c)
	if err != nil {
		logger.Log().Error(err)
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		logger.Log().Error(err)
		return nil, err
	}

	jwtToken, err := auth.GenerateJWTToken(dbUser, auth.SecretKey)
	if err != nil {
		logger.Log().Error(err)
		return nil, err
	}
	expiresIn := time.Hour * time.Duration(1)

	token := &model.Token{
		AccessToken: jwtToken,
		TokenType:   "bearer",
		ExpiresIN:   int(expiresIn.Seconds()),
	}

	return token, nil
}

func (uc *userUseCase) Register(c *gin.Context) error {
	user := &model.User{}
	err := c.BindJSON(user)
	if err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hash)

	return nil
}

func (uc *userUseCase) Take(c *gin.Context) (*model.User, error) {

	user := &model.User{}
	err := c.ShouldBindBodyWith(user, binding.JSON)
	if err != nil {
		logger.Log().Error(err)
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte("test123"), bcrypt.DefaultCost)
	if err != nil {
		logger.Log().Error(err)
		return nil, err
	}

	return &model.User{
		Username: "harvey",
		Password: string(hash),
	}, nil
}
