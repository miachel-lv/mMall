package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
	"user/model"
)

var DB *gorm.DB

func Database(connString string) error {
	fmt.Println("database init, config: ", connString)
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		return err
	}

	db.LogMode(true)
	//if gin.Mode() == "release" {
	//	db.LogMode(false)
	//}

	//默认不加复数
	db.SingularTable(true)
	//设置空闲池
	//空闲
	db.DB().SetMaxIdleConns(20)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	migration()

	return nil
}

func migration() {
	DB.Set(`gorm:table_option`, "charset=utf8mb4").AutoMigrate(&model.User{})
}

//func NewDBClient(ctx context.Context) *gorm.DB {
//	db := DB
//	return db.WithContext(ctx)
//}