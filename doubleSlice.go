package main

import "fmt"

// 接收整数切片指针，每个元素乘以2
func doubleSlice(slicePtr *[]int) {
    // 解引用获取原始切片
    s := *slicePtr
    
    // 遍历每个元素并修改值
    for i := range s {
        s[i] = s[i] * 2
    }
}

func main() {
    // 创建初始切片
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Println("原始切片:", numbers)
    
    // 传递切片指针
    doubleSlice(&numbers)
    
    fmt.Println("修改后切片:", numbers)
    // 输出: [2 4 6 8 10]
}
