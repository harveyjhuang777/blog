package user

import (
	"github.com/jinzhu/gorm"
	"github.com/jwjhuang/blog/service/model"
)

func newProfileDAO() IProfileDAO {
	return &profileDAO{}
}

type profileDAO struct{}

//IProfileDAO is responsible for profile data access controll
type IProfileDAO interface {
	//IsFollowing help to know one user is following another user
	IsFollowing(db *gorm.DB, followed, following uint) bool
	FollowUser(db *gorm.DB, followed, following uint) error
	UnFollowUser(db *gorm.DB, followed, following uint) error
}

func (dao *profileDAO) IsFollowing(db *gorm.DB, followed, following uint) bool {
	var follow model.Follow

	db.Where(model.Follow{
		FollowingID:  following,
		FollowedByID: followed,
	}).First(&follow)

	return follow.ID != 0
}

func (dao *profileDAO) FollowUser(db *gorm.DB, followed, following uint) error {
	follow := &model.Follow{
		FollowingID:  following,
		FollowedByID: followed,
	}

	if err := db.Create(follow).Error; err != nil {
		return err
	}

	return nil
}

func (dao *profileDAO) UnFollowUser(db *gorm.DB, followed, following uint) error {
	follow := &model.Follow{
		FollowingID:  following,
		FollowedByID: followed,
	}

	if err := db.Delete(follow).Error; err != nil {
		return err
	}

	return nil
}
