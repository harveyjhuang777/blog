package user

import (
	"github.com/jinzhu/gorm"
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
	Insert(db *gorm.DB, data *model.User) error
	GetUserByEmail(db *gorm.DB, email string) (*model.User, error)
}

func (md *userDAO) Insert(db *gorm.DB, data *model.User) error {
	if err := db.Create(data).Error; err != nil {
		return err
	}
	return nil
}

func (md *userDAO) GetUserByEmail(db *gorm.DB, email string) (*model.User, error) {
	user := &model.User{}

	if err := db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
