package user

type UserRepository interface {
	FindAll() []*User
	FindUserOfId(id int) (*User, error)
}
