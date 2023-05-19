//go:build wireinject

package wire

import (
	"example-wire/onion_architecture/controller"
	"example-wire/onion_architecture/repository"
	"example-wire/onion_architecture/service"
	"github.com/google/wire"
)

func InitControllers() controller.Controller {
	wire.Build(controller.NewControllers, service.NewServices, repository.NewRepositories)
	return nil
}
