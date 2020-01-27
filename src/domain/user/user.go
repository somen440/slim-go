package user

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewUser(id int, name string) *User {
	return &User{
		ID:   id,
		Name: name,
	}
}
