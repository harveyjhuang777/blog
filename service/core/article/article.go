package article

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jwjhuang/blog/service/app/logger"
	"github.com/jwjhuang/blog/service/model"
	"github.com/jwjhuang/blog/service/utils/errs"
)

func newArticle() IArticleCenter {
	return &articleUseCase{}
}

//IArticleCenter define article's capabilities
type IArticleCenter interface {
	Create(c *gin.Context, cond *model.ArticleCreateCond) (*model.Article, error)
	List(c *gin.Context, cond *model.ArticleGetCond) ([]*model.Article, error)
	GetBySlug(c *gin.Context, slug string) (*model.Article, error)
}

type articleUseCase struct {
}

func (uc *articleUseCase) Create(c *gin.Context, cond *model.ArticleCreateCond) (*model.Article, error) {
	article, err := validateCreateArticle(c, cond)
	if err != nil {
		logger.Log().Error(err)
		return nil, err
	}

	id, err := dao.Article.Insert(packet.DB, article)
	if err != nil {
		logger.Log().Error(err)
		return nil, err
	}

	res, err := dao.Article.GetByID(packet.DB, id)
	if err != nil {
		logger.Log().Error(err)
		return nil, err
	}
	return res, nil
}

func (uc *articleUseCase) List(c *gin.Context, cond *model.ArticleGetCond) ([]*model.Article, error) {
	query := model.NewQueryCond(cond)
	res, err := dao.Article.List(packet.DB, query)
	if err != nil {
		logger.Log().Error(err)
		return nil, err
	}
	return res, nil
}

func (uc *articleUseCase) GetBySlug(c *gin.Context, slug string) (*model.Article, error) {
	res, err := dao.Article.GetBySlug(packet.DB, slug)
	if err != nil {
		logger.Log().Error(err)
		return nil, err
	}
	return res, nil
}

func validateCreateArticle(c *gin.Context, cond *model.ArticleCreateCond) (*model.Article, error) {
	if cond.Title == nil {
		return nil, errs.ErrFieldNotExist
	}

	if cond.Description == nil {
		return nil, errs.ErrFieldNotExist
	}

	if cond.Body == nil {
		return nil, errs.ErrFieldNotExist
	}

	resp := &model.Article{
		Title:       *cond.Title,
		Description: *cond.Description,
		Body:        *cond.Body,
	}

	if len(cond.TagList) > 0 {
		tagList := []*model.Tag{}
		for _, t := range cond.TagList {
			tag := &model.Tag{
				Tag: t,
			}
			tagList = append(tagList, tag)
		}
		resp.Tags = tagList
	}

	resp.Slug = strings.Replace(resp.Title, " ", "-", -1)

	return resp, nil
}
