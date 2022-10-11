package core

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	"user/model"
	"user/services"
)

type UserService struct {

}

func (*UserService) UserLogin(ctx context.Context, req *services.UserRequest, resp *services.UserDetailResponse) error {
	var user model.User
	resp.Code = 200
	if err := model.DB.Where("user_name=?", req.UserName).First(&user).Error; err != nil {
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
	if req.Password != req.PasswordConfirm {
		err := errors.New("两次密码输入不一致")
		return err
	}

	count := 0
	if err := model.DB.Model(&model.User{}).Where("user_name=?", req.UserName).Count(&count).Error; err != nil {
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

	if err := model.DB.Create(&user).Error; err != nil {
		return err
	}

	resp.UserDetail = user.ToUserModel()
	return nil
}