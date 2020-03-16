package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jwjhuang/blog/service/app/logger"
	"github.com/jwjhuang/blog/service/core/user"
	"github.com/jwjhuang/blog/service/model"
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
	user := &model.User{}
	err := c.ShouldBindBodyWith(user, binding.JSON)
	if err != nil {
		logger.Log().Error(err)
		abortWithError(c, http.StatusBadRequest, err)
		return
	}

	res, err := uc.core.Login(c, user)
	if err != nil {
		logger.Log().Error(err)
		abortWithError(c, http.StatusUnauthorized, err)
		return
	}
	responseWithJSON(c, res)
}

func (uc *userController) Register(c *gin.Context) {
	user := &model.User{}
	err := c.BindJSON(user)
	if err != nil {
		logger.Log().Error(err)
		abortWithError(c, http.StatusBadRequest, err)
		return
	}

	if err := uc.core.Register(c, user); err != nil {
		logger.Log().Error(err)
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	responseWithJSON(c, nil)
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
