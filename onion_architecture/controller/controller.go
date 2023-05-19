package controller

import "example-wire/onion_architecture/service"

type Controller interface {
	User() User
}

type controller struct {
	user User
}

// NewControllers returns controllers with all required dependencies
func NewControllers(s service.Services) Controller {
	return &controller{
		user: newUser(s),
	}
}

func (c controller) User() User {
	return c.user
}
