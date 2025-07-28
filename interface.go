package main

import (
	"fmt"
	"math"
)

// Shape 接口定义
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle 结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// Rectangle的方法实现
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle 结构体
type Circle struct {
	Radius float64
}

// Circle的方法实现
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 打印形状信息的辅助函数
func printShapeInfo(s Shape) {
	// 使用类型断言获取具体类型
	switch shape := s.(type) {
	case Rectangle:
		fmt.Printf("矩形: %.2f×%.2f\n", shape.Width, shape.Height)
	case Circle:
		fmt.Printf("圆形: 半径%.2f\n", shape.Radius)
	}
	
	// 调用接口方法
	fmt.Printf("面积: %.2f\n", s.Area())
	fmt.Printf("周长: %.2f\n", s.Perimeter())
	fmt.Println()
}

func main() {
	// 创建Rectangle实例
	rect := Rectangle{Width: 5, Height: 7}
	fmt.Println("矩形信息:")
	printShapeInfo(rect)
	
	// 创建Circle实例
	circle := Circle{Radius: 3}
	fmt.Println("圆形信息:")
	printShapeInfo(circle)
	
	// 通过接口变量调用方法
	var s Shape
	s = rect
	fmt.Println("通过Shape接口访问矩形:")
	fmt.Printf("面积: %.2f\n", s.Area())
	
	// 创建一个Shape切片
	fmt.Println("\n形状集合:")
	shapes := []Shape{
		Rectangle{Width: 4, Height: 6},
		Circle{Radius: 2.5},
		Rectangle{Width: 3, Height: 3},
	}
	
	for i, shape := range shapes {
		fmt.Printf("形状 #%d: ", i+1)
		fmt.Printf("面积 %.2f, 周长 %.2f\n", shape.Area(), shape.Perimeter())
	}
}
