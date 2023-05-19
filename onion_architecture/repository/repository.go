package repository

type Repositories interface {
	User() User
}

type repository struct {
	user User
}

func NewRepositories() Repositories {
	return &repository{
		user: newUser(),
	}
}

func (r repository) User() User {
	return r.user
}
