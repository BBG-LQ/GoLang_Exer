package main

import (
	"BLOG/controller"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {

	// 创建数据库连接
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test001?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	// db.AutoMigrate(&models.Users{}, &models.Posts{}, &models.Comments{}, models.Token{})

	defer db.Close()

	fmt.Println("123")

	r := gin.Default()
	controller.UserControllerInit(r)
	controller.PostControllerInit(r)
	controller.CommentControllerInit(r)
	r.Run(":8080")
}
