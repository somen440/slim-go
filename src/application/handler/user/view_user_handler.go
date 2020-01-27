package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/somen440/slim-go/src/domain/user"
)

type ViewUserHandler struct {
	repository user.UserRepository
}

func NewViewUserHandler(repository user.UserRepository) *ViewUserHandler {
	return &ViewUserHandler{
		repository: repository,
	}
}

func (h *ViewUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	os.Stdout.Write([]byte("ViewUserHandler ServeHTTP\n"))

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	os.Stdout.Write([]byte(fmt.Sprintf("  id: [%d]\n", id)))

	response, err := h.repository.FindUserOfId(id)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	os.Stdout.Write([]byte(fmt.Sprintf("  user: [%v]\n", response)))

	res, err := json.Marshal(response)
	if err != nil {
		fmt.Fprintf(w, string("error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(res))
}
