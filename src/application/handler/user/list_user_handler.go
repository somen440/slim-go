package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/somen440/slim-go/src/domain/user"
)

type ListUserHandler struct {
	repository user.UserRepository
}

func NewListUserHandler(repository user.UserRepository) *ListUserHandler {
	return &ListUserHandler{
		repository: repository,
	}
}

func (h *ListUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	os.Stdout.Write([]byte("ListUserHandler ServeHTTP\n"))

	response := h.repository.FindAll()
	res, err := json.Marshal(response)
	if err != nil {
		fmt.Fprintf(w, string("error"))
		return
	}
	os.Stdout.Write([]byte(fmt.Sprintf("  users: [%v]\n", string(res))))

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(res))
}
