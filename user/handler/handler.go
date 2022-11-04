package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/asim/go-micro/v3"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"user/dao"
	"user/model"
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

func (*UserService) UserLogin(ctx context.Context, req *services.UserRequest, resp *services.UserDetailResponse) error {
	fmt.Println("----UserLogin")
	var user model.User
	resp.Code = 200
	if err := dao.DB.Where("user_name=?", req.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			resp.Code = 400
			return nil
		}
	}

	if user.CheckPassword(req.Password) == false {
		resp.Code = 400
		return nil
	}

	resp.UserDetail = user.ToUserModel()
	return nil
}

func (*UserService) UserRegister(ctx context.Context, req *services.UserRequest, resp *services.UserDetailResponse) error {
	fmt.Println("----UserRegister")
	if req.Password != req.PasswordConfirm {
		err := errors.New("两次密码输入不一致")
		return err
	}

	count := 0
	if err := dao.DB.Model(&model.User{}).Where("user_name=?", req.UserName).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		err := errors.New("用户名已经存在")
		return err
	}

	user := model.User{
		UserName: req.UserName,
	}
	//密码加密
	if err := user.SetPassword(req.Password); err != nil {
		return err
	}

	if err := dao.DB.Create(&user).Error; err != nil {
		return err
	}

	resp.UserDetail = user.ToUserModel()
	return nil
}


