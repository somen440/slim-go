package user

import (
	"reflect"
	"testing"

	model "github.com/somen440/slim-go/src/domain/user"
	repository "github.com/somen440/slim-go/src/infrastructure/persistence/user"
)

func createUsersProvider() []*model.User {
	return []*model.User{
		&model.User{
			ID:   1,
			Name: "hoge",
		},
		&model.User{
			ID:   2,
			Name: "foo",
		},
		&model.User{
			ID:   3,
			Name: "bar",
		},
	}
}

func TestNewInMemoryUserRepository(t *testing.T) {
	expectUsers := createUsersProvider()
	r := repository.NewInMemoryUserRepository(expectUsers)
	if !reflect.DeepEqual(expectUsers, r.FindAll()) {
		t.Errorf("%v != %v", r.FindAll(), expectUsers)
	}
}

func TestFindUserOfId(t *testing.T) {
	r := repository.NewInMemoryUserRepository(createUsersProvider())

	expectUser := &model.User{
		ID:   2,
		Name: "foo",
	}
	actualUser, err := r.FindUserOfId(2)
	if err != nil {
		t.Errorf("%v", err)
	}
	if !reflect.DeepEqual(expectUser, actualUser) {
		t.Errorf("%v != %v", expectUser, actualUser)
	}
}

func TestFindUserOfIdWithNotFound(t *testing.T) {
	r := repository.NewInMemoryUserRepository(createUsersProvider())

	actualUser, err := r.FindUserOfId(len(r.FindAll()) + 1)
	if err == nil {
		t.Errorf("%v", err)
	}
	if actualUser != nil {
		t.Errorf("%v", actualUser)
	}
}
