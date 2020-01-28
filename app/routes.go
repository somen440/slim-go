package app

import (
	"fmt"
	"net/http"

	handler "github.com/somen440/slim-go/src/application/handler/user"
	model "github.com/somen440/slim-go/src/domain/user"
	repository "github.com/somen440/slim-go/src/infrastructure/persistence/user"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	r := mux.NewRouter()

	repository := repository.NewInMemoryUserRepository([]*model.User{
		model.NewUser(1, "hoge"),
		model.NewUser(2, "foo"),
		model.NewUser(3, "bar"),
	})

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	usersRouter := r.PathPrefix("/users").Subrouter()
	usersRouter.Handle("", handler.NewListUserHandler(repository))
	usersRouter.Handle("/{id}", handler.NewViewUserHandler(repository))

	return r
}
