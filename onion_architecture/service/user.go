package service

import "example-wire/onion_architecture/repository"

type User interface {
	Get() (string, error)
}

type user struct {
	repository repository.Repositories
}

func newUser(r repository.Repositories) User {
	return &user{
		repository: r,
	}
}

func (u user) Get() (string, error) {
	return u.repository.User().Get()
}
