package article

import (
	"sync"

	"github.com/jinzhu/gorm"
	"go.uber.org/dig"
)

var (
	packet articleCenterSet
)

var (
	once sync.Once
	ptr  *articleCenter
)

//NewArticleCenter generate article instance for external usage
func NewArticleCenter(set articleCenterSet) articleCenter {
	once.Do(func() {
		packet = set

		newDAO()

		ptr = &articleCenter{
			Article: newArticle(),
		}
	})

	return *ptr
}

type articleCenterSet struct {
	dig.In

	DB *gorm.DB
}

type articleCenter struct {
	dig.Out

	Article IArticleCenter
}
