package main

import (
	"log"

	"github.com/Coreychen4444/Users/model"
	"github.com/Coreychen4444/Users/routers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 连接数据库(用户名和密码自己改)
	dsn := "root:44447777@tcp(:3306)/users?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	// 自动迁移
	db.AutoMigrate(&model.User{})
	// 注册路由
	r := routers.InitRouter(db)
	r.Run(":8080")
}
