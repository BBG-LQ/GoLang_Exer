package main

import (
	"fmt"
	"math"
)

/**
	定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
	在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
**/

type Shape interface {
	Area() float64
	Perimeter() float64
}
type Rectangle struct {
	Width, High float64
}
type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.High

}
func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.High) * 2

}

func (r Circle) Area() float64 {
	return r.Radius * r.Radius * math.Pi

}
func (r Circle) Perimeter() float64 {
	return 2 * math.Pi * r.Radius

}

func main() {
	rect := Rectangle{
		Width: 4,
		High:  3,
	}
	fmt.Println(rect.Area(), rect.Perimeter())

	// cir := Circle{
	// 	Radius: 5,
	// }
	cir := new(Circle)
	cir.Radius = 5
	fmt.Println(cir.Area(), cir.Perimeter())

}
