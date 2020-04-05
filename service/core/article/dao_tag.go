package article

import (
	"github.com/jinzhu/gorm"
	"github.com/jwjhuang/blog/service/app/logger"
	"github.com/jwjhuang/blog/service/model"
)

func newTagDAO() ITagDAO {
	return &tagDAO{}
}

type tagDAO struct{}

//ITagDAO is responsible for tag data access controll
type ITagDAO interface {
	Insert(db *gorm.DB, data *model.Tag) error
	InsertMany(db *gorm.DB, data []*model.Tag) error
}

func (uc *tagDAO) Insert(db *gorm.DB, data *model.Tag) error {
	if err := db.Create(data).Error; err != nil {
		logger.Log().Error(err)
		return err
	}

	return nil
}

func (uc *tagDAO) InsertMany(db *gorm.DB, data []*model.Tag) error {
	if err := db.Create(data).Error; err != nil {
		logger.Log().Error(err)
		return err
	}

	return nil
}
