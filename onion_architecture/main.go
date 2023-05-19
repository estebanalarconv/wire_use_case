package main

import (
	"example-wire/onion_architecture/wire"
	"fmt"
)

func main() {
	//it is a normal dependency injection
	/*
		repositories := repository.NewRepositories()
		services := service.NewServices(repositories)
		controllers := controller.NewControllers(services)
		fmt.Println(controllers.User().Get())*/

	//with wire library
	controllers := wire.InitControllers()
	fmt.Println(controllers.User().Get())
}
