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
	UserName string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm"`
}

func (self *UserService) Verify() bool {
	if self.UserName == "" {
		return false
	}
	if self.Password == "" {
		return false
	}
	return true
}

//var service micro.Service
var cl services.UserService

func init() {
	service := micro.NewService()
	service.Init()
	cl = services.NewUserService("user", service.Client())
}

func (self *UserService)LoginHandler(c *gin.Context) serializer.Response {
	if !self.Verify() {
		return serializer.Response{
			Status: e.InvalidParams,
			//TODO 添加 Status code对应的string.
			//Msg:
		}
	}

	fmt.Println("username: ", self.UserName, " password: ", self.Password)
	resp, err := cl.UserLogin(context.Background(), &services.UserRequest{
		UserName: self.UserName,
		Password: self.Password,
	})
	if err != nil {
		fmt.Println(err)
		return serializer.Response{
			Status: e.ERROR,
			//Msg:
		}
	}

	fmt.Println("UserDetail: ", resp.UserDetail.String())

	return serializer.Response{
		Status: e.SUCCESS,
		Msg: resp.UserDetail.String(),
	}
}

func (self *UserService)RegisterHandler() serializer.Response {
	if !self.Verify() {
		return serializer.Response{
			Status: e.InvalidParams,
			//TODO 添加 Status code对应的string.
			//Msg:
		}
	}

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


