package user

type UserNotFoundException struct{}

func NewUserNotFoundException() *UserNotFoundException {
	return &UserNotFoundException{}
}

func (e *UserNotFoundException) Error() string {
	return "The user you requested does not exist."
}
