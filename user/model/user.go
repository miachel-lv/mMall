package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"user/services"
)

const (
	PasswordCost = 12
)

type User struct {
	gorm.Model
	UserName string `gorm:"unique"`
	PassWordDigest string
}

//加密密码
func(user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}

	user.PassWordDigest = string(bytes)
	return nil
}

//检查密码
func(user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PassWordDigest), []byte(password))
	return nil == err
}

func(user *User) ToUserModel() *services.UserModel {
	userModel := services.UserModel{
		ID: uint32(user.ID),
		UserName: user.UserName,
		CreateAt: user.CreatedAt.Unix(),
		UpdateAt: user.UpdatedAt.Unix(),
	}
	return &userModel
}