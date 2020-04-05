package article

var (
	dao *storage
)

func newDAO() {
	dao = &storage{
		Article: newArticleDAO(),
		Tag:     newTagDAO(),
	}
}

type storage struct {
	Article IArticleDAO
	Tag     ITagDAO
}
