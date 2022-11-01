package handler

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/v3"
	"github.com/gin-gonic/gin"
	"user/pkg/e"
	"user/serializer"
	"user/services"
)

type UserService struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

func (self *UserService) Verify() bool {
	//TODO
	return false
}

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

	c.JSON(e.SUCCESS, gin.H{"code": e.SUCCESS, "msg": resp.UserDetail})
}

func (self *UserService)RegisterHandler() serializer.Response {
	if !self.Verify() {
		return serializer.Response{
			Status: e.InvalidParams,
			//TODO 添加 Status code对应的string.
			//Msg:
		}
	}

	cl := services.NewUserService("user", service.Client())
	resp, err := cl.UserRegister(context.Background(), &services.UserRequest{
		UserName: self.UserName,
		Password: self.Password,
		PasswordConfirm: self.PasswordConfirm,
	})
	if err != nil {
		fmt.Println("micro UserRegister fail: ", err.Error())
		//TODO 区分micro 返回的错误码
		return serializer.Response{
			Status: e.ERROR,
			//Msg:
		}
	}

	return serializer.Response{
		Status: e.SUCCESS,
		Msg: resp.String(),
	}
}


