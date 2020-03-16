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
	Register(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
	GetProfile(c *gin.Context)
	Follow(c *gin.Context)
	UnFollow(c *gin.Context)
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

func (uc *userController) Register(c *gin.Context) {
}

func (uc *userController) Get(c *gin.Context) {
}

func (uc *userController) Update(c *gin.Context) {
}

func (uc *userController) GetProfile(c *gin.Context) {
}

func (uc *userController) Follow(c *gin.Context) {
}

func (uc *userController) UnFollow(c *gin.Context) {
}
