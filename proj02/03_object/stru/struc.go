package main

import "fmt"

/**

	使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
	为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。

**/

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Person
	EmployeeID int
}

func (e Employee) PrintInfo() {
	fmt.Printf("姓名： %s\n年龄：%d\n员工ID：%d", e.Name, e.Age, e.EmployeeID)

}

func main() {
	emp := Employee{
		EmployeeID: 123,
		Person: Person{
			Name: "lq",
			Age:  19,
		},
	}
	emp.PrintInfo()

}
