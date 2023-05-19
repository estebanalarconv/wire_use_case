package controller

import "example-wire/onion_architecture/service"

type User interface {
	Get() (string, error)
}

type user struct {
	service service.Services
}

func newUser(s service.Services) *user {
	return &user{service: s}
}

func (u user) Get() (string, error) {
	return u.service.User().Get()
}
