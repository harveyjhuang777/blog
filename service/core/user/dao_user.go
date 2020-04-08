package user

import (
	"github.com/jinzhu/gorm"
	"github.com/jwjhuang/blog/service/model"
)

func newUserDAO() IUserDAO {
	return &userDAO{}
}

type userDAO struct{}

//IUserDAO is responsible for user data access controll
type IUserDAO interface {
	Insert(db *gorm.DB, data *model.User) error
	Update(db *gorm.DB, data *model.User) error
	GetUserByEmail(db *gorm.DB, email string) (*model.User, error)
	GetUserByCondition(db *gorm.DB, cond model.IQueryCond) (*model.User, error)
}

func (ud *userDAO) Insert(db *gorm.DB, data *model.User) error {
	if err := db.Create(data).Error; err != nil {
		return err
	}
	return nil
}

func (ud *userDAO) Update(db *gorm.DB, data *model.User) error {
	if err := db.Save(data).Error; err != nil {
		return err
	}
	return nil
}

func (ud *userDAO) GetUserByEmail(db *gorm.DB, email string) (*model.User, error) {
	user := &model.User{}

	if err := db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ud *userDAO) GetUserByCondition(db *gorm.DB, cond model.IQueryCond) (*model.User, error) {
	user := &model.User{}

	if err := db.Where(cond.Where()).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
