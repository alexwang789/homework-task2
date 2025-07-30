package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	count int
	mu    sync.Mutex
	wg    sync.WaitGroup
}

func (sc *SafeCounter) Increment() {
	defer sc.wg.Done()
	
	for i := 0; i < 1000; i++ {
		sc.mu.Lock()    // 获取锁
		sc.count++     // 安全递增
		sc.mu.Unlock()  // 释放锁
		
		// 模拟工作负载
		time.Sleep(time.Microsecond)
	}
}

func main() {
	// 创建带互斥锁的计数器
	counter := SafeCounter{
		count: 0,
	}
	
	const numGoroutines = 10
	counter.wg.Add(numGoroutines)
	
	start := time.Now()
	
	fmt.Println("启动", numGoroutines, "个协程进行计数...")
	fmt.Println("---------------------------------")
	
	// 启动10个协程
	for i := 0; i < numGoroutines; i++ {
		go counter.Increment()
	}
	
	// 等待所有协程完成
	counter.wg.Wait()
	duration := time.Since(start)
	
	fmt.Println("---------------------------------")
	fmt.Printf("预期结果: 10000\n实际结果: %d\n", counter.count)
	fmt.Printf("耗时: %v\n", duration)
	
	// 验证结果
	if counter.count == 10000 {
		fmt.Println("✅ 测试通过: 结果正确!")
	} else {
		fmt.Printf("❌ 测试失败: 期望10000，实际%d\n", counter.count)
	}
}
