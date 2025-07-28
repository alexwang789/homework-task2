package main

import (
	"fmt"
	"sync"
	"time"
)

func printOdds(wg *sync.WaitGroup) {
	defer wg.Done() // 协程完成时通知WaitGroup
	for i := 1; i <= 10; i += 2 {
		fmt.Printf("奇数: %d\n", i)
		time.Sleep(100 * time.Millisecond) // 添加延迟观察并发效果
	}
}

func printEvens(wg *sync.WaitGroup) {
	defer wg.Done() // 协程完成时通知WaitGroup
	for i := 2; i <= 10; i += 2 {
		fmt.Printf("偶数: %d\n", i)
		time.Sleep(100 * time.Millisecond) // 添加延迟观察并发效果
	}
}

func main() {
	var wg sync.WaitGroup // 使用WaitGroup同步协程
	
	wg.Add(2) // 添加2个协程的计数
	
	// 启动奇数打印协程
	go printOdds(&wg)
	
	// 启动偶数打印协程
	go printEvens(&wg)
	
	wg.Wait() // 等待所有协程完成
	fmt.Println("所有协程已完成!")
}
