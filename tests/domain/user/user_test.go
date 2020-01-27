package user

import (
	"reflect"
	"testing"

	"github.com/somen440/slim-go/src/domain/user"
)

func TestNewUser(t *testing.T) {
	expectID := 123
	expectName := "abcdefg"

	createdUser := &user.User{
		ID:   expectID,
		Name: expectName,
	}

	if expectID != createdUser.ID {
		t.Errorf("%v != %v", expectID, createdUser.ID)
	}

	if expectName != createdUser.Name {
		t.Errorf("%v != %v", expectName, createdUser.Name)
	}

	newUser := user.NewUser(expectID, expectName)
	if !reflect.DeepEqual(createdUser, newUser) {
		t.Errorf("%v != %v", createdUser, newUser)
	}
}
