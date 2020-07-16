package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func main() {
	//打开数据库连接
	db, _ := gorm.Open("mysql", "root:root@(localhost)/gorm-demo?charset=utf8&parseTime=True&loc=Local")

	//关闭数据库连接
	defer db.Close()
}