package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"user/conf"
	"user/core"
	"user/services"
)

func main() {
	conf.Init()

	srv := service.New(
		service.Name("user"),
		//service.Address("127.0.0.1:8001"),
	)

	srv.Init()
	// Register handler
	services.RegisterUserServiceHandler(srv.Server(), new(core.UserService))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
