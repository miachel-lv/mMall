package dao

import "github.com/jinzhu/gorm"

type UserDao struct {
	*gorm.DB
}

//func NewUserDao(ctx context.Context) {
//	return &UserDao{}
//}