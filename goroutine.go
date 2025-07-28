package main

import (
	"fmt"
	"sync"
)

func printOdds(wg *sync.WaitGroup, ch chan struct{}) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		<-ch // 等待信号
		fmt.Printf("线程1(奇数): %d\n", i)
		ch <- struct{}{} // 通知对方
	}
}

func printEvens(wg *sync.WaitGroup, ch chan struct{}) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		ch <- struct{}{} // 先发出信号
		fmt.Printf("线程2(偶数): %d\n", i)
		<-ch // 等待对方完成
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan struct{}, 1) // 缓冲通道控制顺序
	
	wg.Add(2)
	go printOdds(&wg, ch)
	go printEvens(&wg, ch)
	
	wg.Wait()
	close(ch)
}
