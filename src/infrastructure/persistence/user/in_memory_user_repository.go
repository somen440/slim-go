package user

import (
	"github.com/somen440/slim-go/src/domain/user"
)

type InMemoryUserRepository struct {
	users []*user.User
}

func NewInMemoryUserRepository(users []*user.User) *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: users,
	}
}

func (r *InMemoryUserRepository) FindAll() []*user.User {
	return r.users
}

func (r *InMemoryUserRepository) FindUserOfId(id int) (*user.User, error) {
	for _, v := range r.users {
		if v.ID == id {
			return v, nil
		}
	}
	return nil, user.NewUserNotFoundException()
}
