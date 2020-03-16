package user

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jwjhuang/blog/service/app/logger"
	"github.com/jwjhuang/blog/service/model"
	"github.com/jwjhuang/blog/service/utils/auth"
	"golang.org/x/crypto/bcrypt"
)

func newUser() IUserCenter {
	return &userUseCase{}
}

type IUserCenter interface {
	Login(c *gin.Context, user *model.User) (*model.Token, error)
	Register(c *gin.Context, user *model.User) error
	GetUserByEmail(c *gin.Context, email string) (*model.User, error)
}

type userUseCase struct {
}

func (uc *userUseCase) Login(c *gin.Context, user *model.User) (*model.Token, error) {

	dbUser, err := uc.GetUserByEmail(c, user.Email)
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

func (uc *userUseCase) Register(c *gin.Context, user *model.User) error {

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hash)

	if err := dao.User.Insert(packet.DB, user); err != nil {
		return err
	}

	return nil
}

func (uc *userUseCase) GetUserByEmail(c *gin.Context, email string) (*model.User, error) {
	resp, err := dao.User.GetUserByEmail(packet.DB, email)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
