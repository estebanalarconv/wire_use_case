package repository

type User interface {
	Get() (string, error)
}

type user struct{}

func newUser() User {
	return &user{}
}

func (u user) Get() (string, error) {
	return "Getting user from DB", nil
}
