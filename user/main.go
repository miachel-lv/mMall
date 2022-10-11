package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"user/core"
	"user/services"
)

func main() {
	srv := service.New(
		service.Name("rpcUserService"),
		//service.Address("127.0.0.1:8001"),
	)

	// Register handler
	services.RegisterUserServiceHandler(srv.Server(), new(core.UserService))

	srv.Init()

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
