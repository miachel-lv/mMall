package dao_test_test

import (
	"testing"
	"user/dao"
)

func TestDatabase(t *testing.T) {
	connStr := "root:root1234@tcp(127.0.0.1:3306)/m_mall?charset=utf8&parseTime=True&loc=Local"
	err := dao.Database(connStr)
	if err != nil {
		t.Fatal(err)
	}
}
