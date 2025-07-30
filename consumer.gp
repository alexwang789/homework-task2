package main

import (
	"fmt"
	"sync"
)

func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	// 生产100个整数
	for i := 1; i <= 100; i++ {
		ch <- i // 发送整数到通道
		fmt.Printf("生产者: 已发送 %d (缓冲区: %d/%d)\n", i, len(ch), cap(ch))
	}

	close(ch) // 生产完成后关闭通道
	fmt.Println("生产者: 已完成100个整数的发送")
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	// 持续接收直到通道关闭
	count := 0
	for num := range ch {
		count++
		fmt.Printf("消费者: 接收并处理 %d (已处理: %d/100)\n", num, count)
	}

	fmt.Println("消费者: 通道已关闭，共接收", count, "个整数")
}

func main() {
	const bufferSize = 10  // 缓冲区大小
	const totalItems = 100 // 总生产数量

	// 创建带缓冲的通道（容量为bufferSize）
	ch := make(chan int, bufferSize)

	var wg sync.WaitGroup
	wg.Add(2) // 等待两个协程

	fmt.Printf("启动生产者-消费者模型 (缓冲区大小: %d, 总数量: %d)\n", bufferSize, totalItems)
	fmt.Println("==========================================")

	// 启动生产者协程
	go producer(ch, &wg)

	// 启动消费者协程
	go consumer(ch, &wg)

	// 等待所有协程完成
	wg.Wait()

	fmt.Println("==========================================")
	fmt.Println("主程序: 生产消费任务已完成")
}
