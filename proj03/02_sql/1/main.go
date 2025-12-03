package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/*
*

	假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
	要求 ：
	编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
	编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。

*
*/
/**

	假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
	要求 ：
	定义一个 Book 结构体，包含与 books 表对应的字段。
	编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。

**/

type Employee struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     int    `db:"salary"`
}
type Books struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func createTableSql(db *sqlx.DB) {
	// 创建表
	createTableSql := `
	CREATE TABLE IF NOT EXISTS books(
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(64) NOT NULL,
		author VARCHAR(128) NOT NULL,
		price DECTMAL NOT NULL
	)`
	_, err0 := db.Exec(createTableSql)
	if err0 != nil {
		fmt.Println("创建表失败: ", err0)
		return
	}

}

func insertEmployees(db *sqlx.DB) {

	// employeeList := []*Employee{
	// 	{1, "name1", "技术部", 6000},
	// 	{2, "name2", "技术部", 9000},
	// 	{3, "name3", "技术部", 10000},
	// 	{4, "name4", "销售部", 6000},
	// 	{5, "name5", "销售部", 7000},
	// }

	booksIns := []*Books{
		{1, "title1", "author1", 10.0},
		{2, "title2", "author2", 20.0},
		{3, "title3", "author3", 9.9},
		{4, "title4", "author4", 100.0},
		{5, "title5", "author5", 99999.99},
	}
	// _, er := db.NamedExec("INSERT INTO employe (id, name, department, salary) VALUES (:id, :name, :department, :salary)", employeeList)
	_, er := db.NamedExec("INSERT INTO books (id, title, author, price) VALUES (:id, :title, :author, :price)", booksIns)
	if er != nil {
		fmt.Println("批量插入测试数据异常：", er)
		return
	}
}

func selectFromDepartment(db *sqlx.DB, department string) ([]Employee, error) {
	var employees []Employee

	err1 := db.Select(&employees, "SELECT * FROM employe WHERE department = ?", department)
	if err1 != nil {
		return employees, err1
	}
	return employees, nil
}

func selectMaxSalary(db *sqlx.DB) (Employee, error) {
	var employees Employee

	err1 := db.Get(&employees, "SELECT * FROM employe order by salary desc limit 1")
	if err1 != nil {
		return employees, err1
	}
	return employees, nil
}

func selectBooks(db *sqlx.DB) ([]Books, error) {
	var selecbooks []Books
	err1 := db.Select(&selecbooks, "SELECT * FROM books where price>50")
	if err1 != nil {
		return nil, err1
	}

	return selecbooks, nil

}

func main() {
	// 连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/test001?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	defer db.Close()

	// createTableSql(db)

	// insertEmployees(db)

	employees, err := selectFromDepartment(db, "技术部")
	if err != nil {
		log.Fatalf("查询员工失败: %v", err)
	}

	for _, employee := range employees {
		fmt.Printf("员工信息: %+v\n", employee)
	}

	employees1, err1 := selectMaxSalary(db)
	if err1 != nil {
		log.Fatalf("查询员工失败: %v", err1)
	}

	fmt.Printf("员工信息: %+v\n", employees1)

	selecbooks, err2 := selectBooks(db)
	if err2 != nil {
		log.Fatalf("查询book失败: %v", err2)
	}
	for _, book := range selecbooks {
		fmt.Printf("book信息: %+v\n", book)
	}

}
