package main

import (
	"fmt"
	"sync"
)

func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // 完成时通知WaitGroup
	
	// 生成整数1到10并发送到通道
	for i := 1; i <= 10; i++ {
		ch <- i  // 发送整数到通道
		fmt.Printf("生产者: 已发送 %d\n", i)
	}
	
	close(ch) // 关闭通道（通知消费者停止接收）
	fmt.Println("生产者: 已完成发送")
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // 完成时通知WaitGroup
	
	// 从通道接收所有整数
	for num := range ch {
		fmt.Printf("消费者: 接收并处理 %d\n", num)
	}
	fmt.Println("消费者: 通道已关闭")
}

func main() {
	// 创建带缓冲的通道（容量5）
	ch := make(chan int, 5)
	
	// 创建WaitGroup等待两个线程完成
	var wg sync.WaitGroup
	wg.Add(2)
	
	// 启动生产者线程
	go producer(ch, &wg)
	
	// 启动消费者线程
	go consumer(ch, &wg)
	
	// 等待所有线程完成
	wg.Wait()
	fmt.Println("\n主程序: 所有线程已完成")
}
