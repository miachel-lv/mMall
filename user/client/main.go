package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/v3"
	"user/services"
)

func main() {
	ServiceOne()
}

func ServiceOne() {
	service := micro.NewService()
	service.Init()

	//创建微服务客户端
	cl := services.NewUserService("user", service.Client())
	resp, err := cl.UserLogin(context.Background(), &services.UserRequest{
		UserName: "mike",
		Password: "123456",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("----status code: ", resp.Code, "\n", resp.UserDetail)
}
