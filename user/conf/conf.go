package conf

import (
	"fmt"
	"os"
	"user/dao"
)

func Init() {
	//defer func() {
	//	if recov := recover(); recov != nil {
	//		fmt.Errorf("init panic at: %v", recov)
	//		return
	//	}
	//}()

	//读取本地环境
	//err := godotenv.Load("./conf/.env.example")
	//if err != nil {
	//	panic(err)
	//}

	fmt.Println("get MYSQL_DSN: ", os.Getenv("MYSQL_DSN"))

	//model.Database(os.Getenv("MYSQL_DSN"))
	dao.Database("root:root1234@tcp(127.0.0.1:3306)/m_mall?charset=utf8&parseTime=True&loc=Local")
}