package micro_server

import (
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"user/core"
	"user/services"
)

func Init() {
	srv :=  micro.NewService(
		micro.Name("user"),
		micro.Version("latest"),
		//micro.Address(":8888"),
	)
	srv.Init()
	// Register handler
	if err := services.RegisterUserServiceHandler(srv.Server(), new(core.UserService)); err != nil {
		logger.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
