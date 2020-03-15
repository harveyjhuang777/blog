package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jwjhuang/blog/service/model"
)

var (
	dao       *storage
	tableName = "user"
)

func newDAO() {
	dao = &storage{
		User: newUserDAO(),
	}
}

type storage struct {
	User IUserDAO
}

func newUserDAO() IUserDAO {
	return &userDAO{}
}

type userDAO struct{}

type IUserDAO interface {
	Insert(ctx *gin.Context, doc *model.User) error
}

func (md *userDAO) Insert(ctx *gin.Context, doc *model.User) error {
	return nil
}
