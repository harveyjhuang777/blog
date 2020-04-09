package article

import (
	"github.com/jinzhu/gorm"
	"github.com/jwjhuang/blog/service/app/logger"
	"github.com/jwjhuang/blog/service/model"
)

func newArticleDAO() IArticleDAO {
	return &articleDAO{}
}

type articleDAO struct{}

//IArticleDAO is responsible for article data access controll
type IArticleDAO interface {
	Insert(db *gorm.DB, data *model.Article) (uint, error)
	List(db *gorm.DB, cond model.IQueryCond) ([]*model.Article, error)
	GetByID(db *gorm.DB, id uint) (*model.Article, error)
	GetBySlug(db *gorm.DB, slug string) (*model.Article, error)
	GetByCondition(db *gorm.DB, cond model.IQueryCond) (*model.Article, error)
	//Update(db *gorm.DB, data *model.Article) error
}

func (ad *articleDAO) Insert(db *gorm.DB, data *model.Article) (uint, error) {
	tx := db.Begin()
	if err := tx.Save(data).Error; err != nil {
		logger.Log().Error(err)
		tx.Rollback()
		return data.ID, err
	}

	return data.ID, tx.Commit().Error
}

func (ad *articleDAO) List(db *gorm.DB, query model.IQueryCond) ([]*model.Article, error) {
	var tags []*model.Tag

	res := []*model.Article{}
	paging := query.Paging()
	if err := db.Where(query.Where()).Find(&res).Offset(paging.GetOffset()).Limit(paging.GetSize()).Error; err != nil {
		logger.Log().Error(err)
		return nil, err
	}

	for _, article := range res {
		db.Model(article).Related(&tags, "Tags")
		article.Tags = tags
	}

	return res, nil
}

func (ad *articleDAO) GetByID(db *gorm.DB, id uint) (*model.Article, error) {
	var tags []*model.Tag

	res := &model.Article{}

	if err := db.Where("id = ?", id).First(res).Error; err != nil {
		logger.Log().Error(err)
		return nil, err
	}

	db.Model(res).Related(&tags, "Tags")

	res.Tags = tags

	return res, nil
}

func (ad *articleDAO) GetBySlug(db *gorm.DB, slug string) (*model.Article, error) {
	var tags []*model.Tag

	res := &model.Article{}

	if err := db.Where("slug = ?", slug).First(res).Error; err != nil {
		logger.Log().Error(err)
		return nil, err
	}

	db.Model(res).Related(&tags, "Tags")

	res.Tags = tags

	return res, nil
}

func (ad *articleDAO) GetByCondition(db *gorm.DB, cond model.IQueryCond) (*model.Article, error) {
	var tags []*model.Tag

	res := &model.Article{}

	if err := db.Where(cond.Where()).First(res).Error; err != nil {
		logger.Log().Error(err)
		return nil, err
	}

	db.Model(res).Related(&tags, "Tags")

	res.Tags = tags

	return res, nil
}
