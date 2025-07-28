package main

import "fmt"

// 定义Person结构体
type Person struct {
	Name string
	Age  int
}

// 定义Employee结构体，组合Person
type Employee struct {
	Person      // 匿名字段（组合）
	EmployeeID string
}

// Employee的方法：打印员工信息
func (e Employee) PrintInfo() {
	fmt.Printf("员工信息:\n")
	fmt.Printf("姓名: %s\n", e.Name)        // 直接访问Person字段
	fmt.Printf("年龄: %d\n", e.Age)          // 直接访问Person字段
	fmt.Printf("员工ID: %s\n", e.EmployeeID) // 访问自身字段
}

func main() {
	// 创建Person实例
	person := Person{
		Name: "张三",
		Age:  30,
	}

	// 创建Employee实例（组合Person）
	employee := Employee{
		Person:     person, // 嵌入Person对象
		EmployeeID: "E1001",
	}
	
	// 直接调用Employee的方法
	fmt.Println("使用PrintInfo方法输出:")
	employee.PrintInfo()
	
	// 也可以创建时直接初始化
	manager := Employee{
		Person: Person{
			Name: "李四",
			Age:  40,
		},
		EmployeeID: "E2001",
	}
	
	fmt.Println("\n经理信息:")
	manager.PrintInfo()
	
	// 演示组合的特性：可以直接访问嵌入结构体的字段
	fmt.Println("\n直接访问嵌入字段:")
	employee.Name = "王五" // 直接修改嵌入结构体的字段
	employee.Age = 35
	fmt.Println("修改后姓名:", employee.Name)
	fmt.Println("修改后年龄:", employee.Age)
	
	// 创建Employee切片
	fmt.Println("\n员工列表:")
	employees := []Employee{
		{Person{"赵六", 28}, "E3001"},
		{Person{"钱七", 32}, "E3002"},
		{Person{"孙八", 45}, "E3003"},
	}
	
	for _, emp := range employees {
		emp.PrintInfo()
		fmt.Println("------")
	}
}
