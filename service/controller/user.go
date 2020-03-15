package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jwjhuang/blog/service/core/user"
)

func newUserController(core user.IUserCenter) IUserController {
	return &userController{
		core: core,
	}
}

type IUserController interface {
	Login(c *gin.Context)
}

type userController struct {
	core user.IUserCenter
}

func (uc *userController) Login(c *gin.Context) {
	res, err := uc.core.Login(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}
