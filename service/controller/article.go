package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jwjhuang/blog/service/app/logger"
	"github.com/jwjhuang/blog/service/core/article"
	"github.com/jwjhuang/blog/service/model"
	"github.com/jwjhuang/blog/service/utils/errs"
)

func newArticleController(core article.IArticleCenter) IArticleController {
	return &articleController{
		core: core,
	}
}

type IArticleController interface {
	List(c *gin.Context)
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	TagList(c *gin.Context)
}

type articleController struct {
	core article.IArticleCenter
}

func (uc *articleController) List(c *gin.Context) {
	cond := &model.ArticleGetCond{}
	err := c.ShouldBindQuery(cond)
	if err != nil {
		logger.Log().Error(err)
		abortWithError(c, http.StatusBadRequest, err)
		return
	}

	res, err := uc.core.List(c, cond)
	if err != nil {
		logger.Log().Error(err)
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	responseWithJSON(c, res)
}

func (uc *articleController) Get(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		logger.Log().Error(errs.ErrInvalidRequest)
		abortWithError(c, http.StatusBadRequest, errs.ErrInvalidRequest)
		return
	}

	res, err := uc.core.GetBySlug(c, slug)
	if err != nil {
		logger.Log().Error(err)
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	responseWithJSON(c, res)
}

func (uc *articleController) Create(c *gin.Context) {
	cond := &model.ArticleCreateCond{}
	err := c.ShouldBindBodyWith(cond, binding.JSON)
	if err != nil {
		logger.Log().Error(err)
		abortWithError(c, http.StatusBadRequest, err)
		return
	}

	res, err := uc.core.Create(c, cond)
	if err != nil {
		logger.Log().Error(err)
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	responseWithJSON(c, res)
}

func (uc *articleController) Update(c *gin.Context) {
}

func (uc *articleController) Delete(c *gin.Context) {
}

func (uc *articleController) TagList(c *gin.Context) {
}
