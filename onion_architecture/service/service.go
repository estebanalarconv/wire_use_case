package service

import "example-wire/onion_architecture/repository"

type Services interface {
	User() User
}

type service struct {
	user User
}

func NewServices(r repository.Repositories) Services {
	return &service{
		user: newUser(r),
	}
}

func (s service) User() User {
	return s.user
}
