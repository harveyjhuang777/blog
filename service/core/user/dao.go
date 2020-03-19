package user

var (
	dao *storage
)

func newDAO() {
	dao = &storage{
		User:    newUserDAO(),
		Profile: newProfileDAO(),
	}
}

type storage struct {
	User    IUserDAO
	Profile IProfileDAO
}
