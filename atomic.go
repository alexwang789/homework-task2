package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// 初始化原子计数器
	var counter int64 = 0
	
	const numGoroutines = 10
	const incrementsPerGoroutine = 1000
	
	var wg sync.WaitGroup
	wg.Add(numGoroutines)
	
	start := time.Now()
	
	fmt.Println("启动", numGoroutines, "个协程进行原子计数...")
	fmt.Println("---------------------------------")
	
	// 启动10个协程执行原子递增
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < incrementsPerGoroutine; j++ {
				// 使用原子操作递增计数器
				atomic.AddInt64(&counter, 1)
				
				// 每100次操作报告进度
				if j%100 == 0 {
					fmt.Printf("协程%d: 完成%d次递增\n", id, j)
				}
				
				// 模拟工作负载（非必要）
				time.Sleep(time.Microsecond)
			}
		}(i)
	}
	
	// 中间进度报告（演示原子加载）
	go func() {
		for {
			current := atomic.LoadInt64(&counter)
			if current >= int64(numGoroutines*incrementsPerGoroutine) {
				break
			}
			
			fmt.Printf("进度监控: 当前计数 = %d\n", current)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	
	// 等待所有工作协程完成
	wg.Wait()
	duration := time.Since(start)
	
	// 获取最终计数
	finalCount := atomic.LoadInt64(&counter)
	
	fmt.Println("---------------------------------")
	fmt.Printf("预期结果: %d\n", numGoroutines*incrementsPerGoroutine)
	fmt.Printf("实际结果: %d\n", finalCount)
	fmt.Printf("耗时: %v\n", duration)
	
	// 验证结果
	if finalCount == int64(numGoroutines*incrementsPerGoroutine) {
		fmt.Println("✅ 测试通过: 原子操作正确执行!")
	} else {
		fmt.Printf("❌ 测试失败: 期望%d，实际%d\n", 
			numGoroutines*incrementsPerGoroutine, finalCount)
	}
}
