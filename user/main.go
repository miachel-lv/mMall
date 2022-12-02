package main

import (
	"fmt"
	httpServer "github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"github.com/gin-gonic/gin"
	"user/api"
	"user/conf"
	"user/micro_server"
)

func main() {
	conf.Init()
	go func() {
		fmt.Println("Starting micro server.")
		micro_server.Init()
	}()

	srv := httpServer.NewServer(
		server.Name("go.micro.web.DemoHTTP"),
		server.Address(":8082"),
	)

	router := gin.New()
	router.Use(gin.Recovery())
	router.POST("/user/login", api.UserLogin)
	router.POST("/user/register", api.UserRegister)
	router.GET("/",  api.Index)

	hd := srv.NewHandler(router)
	if err := srv.Handle(hd); err != nil {
		logger.Fatal(err)
	}

	service :=  micro.NewService(
		micro.Server(srv),
		micro.Registry(registry.NewRegistry()),
	)
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}

