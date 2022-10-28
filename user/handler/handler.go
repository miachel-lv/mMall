package handler

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/v3"
	"github.com/gin-gonic/gin"
	"user/services"
)

var service micro.Service

func init() {
	service = micro.NewService()
	service.Init()
}

func LoginHandler(c *gin.Context) {
	name := c.Param("username")
	password := c.Param("password")
	fmt.Println("username: ", name, " password: ", password)

	cl := services.NewUserService("user", service.Client())
	resp, err := cl.UserLogin(context.Background(), &services.UserRequest{
		UserName: name,
		Password: password,
	})
	if err != nil {
		fmt.Println(err)
		c.JSON(201, gin.H{"code": resp.Code, "msg": err.Error()})
		return
	}

	c.JSON(200, gin.H{"code": 200, "msg": resp.UserDetail})
}

func RegisterHandler(c *gin.Context) {

	c.ShouldBind()
	userName := ""
	password := ""

	cl := services.NewUserService("user", service.Client())
	resp, err := cl.UserRegister(context.Background(), &services.UserRequest{
		UserName: userName,
		Password: password,
		PasswordConfirm: password,
	})
	if err != nil {
		fmt.Println("micro UserRegister fail: ", err.Error())
		c.JSON(200, gin.H{"code": resp.Code, "msg": err.Error()})
		return
	}

	c.JSON(200, gin.H{"code": resp.Code, "msg": resp.UserDetail})
}


