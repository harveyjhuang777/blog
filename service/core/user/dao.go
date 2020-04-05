package user

var (
	dao *storage
)

func newDAO() {
	dao = &storage{
		User: newUserDAO(),
	}
}

type storage struct {
	User IUserDAO
}
