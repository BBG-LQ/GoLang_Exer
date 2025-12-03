package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
*

	假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、
	grade （学生年级，字符串类型）。
	要求 ：
	编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。

*
*/
type Student struct {
	ID    uint `gorm:"primaryKey;autoIncrement"`
	Name  string
	Age   int
	Grade string
}

func main() {

	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/test001?charset=utf8mb4&parseTime=True&loc=Local"))

	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	// 创建student表
	db.AutoMigrate(&Student{})

	// 向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"
	student := Student{Name: "张三", Age: 20, Grade: "三年级"}
	result := db.Create(&student)
	if result.Error != nil {
		log.Println("创建失败:", result.Error)
	}
	fmt.Printf("创建成功，ID: %d\n", student.ID)

	// 查询 students 表中所有年龄大于 18 岁的学生信息
	studentList := []Student{}
	result1 := db.Where("Age > ?", 18).Find(&studentList)
	if result1.Error != nil {
		log.Println("查询失败:", result1.Error)
	}
	for _, s := range studentList {
		fmt.Printf(">18学生信息: ID:%d \n", s.ID)
	}

	// 将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"
	result2 := db.Model(&Student{}).Where("Name = ?", "张三").Updates(Student{Grade: "四年级"})
	if result2.Error != nil {
		log.Println("更新失败:", result2.Error)
	}
	fmt.Printf("影响行数 %d\n", result2.RowsAffected)

	// 删除 students 表中年龄小于 15 岁的学生记录
	result3 := db.Where("Age<?", 18).Delete(&Student{})
	if result3 != nil {
		log.Println("删除失败:", result3.Error)
	}
	fmt.Printf("影响行数 %d\n", result3.RowsAffected)

	result4 := db.Where("Id=?", 11).Delete(&Student{})
	if result4.Error != nil {
		log.Println("删除失败:", result4.Error)
	}
	fmt.Printf("影响行数 %d\n", result3.RowsAffected)

}
