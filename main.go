package main

import (
	"log"
	"net/http"
	"time"

	handler "github.com/somen440/slim-go/src/application/handler/user"
	model "github.com/somen440/slim-go/src/domain/user"
	repository "github.com/somen440/slim-go/src/infrastructure/persistence/user"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	repository := repository.NewInMemoryUserRepository([]*model.User{
		model.NewUser(1, "hoge"),
		model.NewUser(2, "foo"),
		model.NewUser(3, "bar"),
	})

	usersRouter := r.PathPrefix("/users").Subrouter()
	usersRouter.Handle("", handler.NewListUserHandler(repository))
	usersRouter.Handle("/{id}", handler.NewViewUserHandler(repository))

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
