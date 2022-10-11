package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

func Database(connString string) {
	fmt.Println("database init, config: ", connString)
	defer func() {
		if recov := recover(); recov != nil {
			fmt.Errorf("panic at: %v", recov)
			return
		}
	}()

	db, err := gorm.Open("mysql", connString)
	if err != nil {
		panic(err)
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
}

func migration() {
	DB.Set(`gorm:table_option`, "charset=utf8mb4").AutoMigrate(&User{})
}