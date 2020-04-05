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
	Find(db *gorm.DB, id uint) (*model.Article, error)
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

func (ad *articleDAO) Find(db *gorm.DB, id uint) (*model.Article, error) {
	res := &model.Article{}

	if err := db.Where("id = ?", id).First(res).Error; err != nil {
		logger.Log().Error(err)
		return nil, err
	}

	return res, nil
}
