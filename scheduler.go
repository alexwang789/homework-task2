package main

import (
	"fmt"
	"sync"
	"time"
)

// Task 定义任务类型
type Task func()

// TaskResult 存储任务执行结果
type TaskResult struct {
	Name     string        // 任务名称
	Duration time.Duration // 执行时间
}

// Scheduler 任务调度器
type Scheduler struct {
	tasks     []Task       // 任务列表
	taskNames []string     // 任务名称列表
	results   []TaskResult // 任务执行结果
	wg        sync.WaitGroup
	startTime time.Time
}

// NewScheduler 创建新的调度器
func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks:   make([]Task, 0),
		results: make([]TaskResult, 0),
	}
}

// AddTask 添加任务
func (s *Scheduler) AddTask(name string, task Task) {
	s.tasks = append(s.tasks, task)
	s.taskNames = append(s.taskNames, name)
}

// worker 工作协程，执行任务
func (s *Scheduler) worker(id int, task Task, name string) {
	defer s.wg.Done()
	
	start := time.Now()
	fmt.Printf("Worker %d 开始执行任务: %s\n", id, name)
	
	task() // 执行任务
	
	duration := time.Since(start)
	s.results = append(s.results, TaskResult{Name: name, Duration: duration})
	fmt.Printf("Worker %d 完成任务: %s (耗时: %v)\n", id, name, duration)
}

// Run 并发执行所有任务
func (s *Scheduler) Run() {
	s.startTime = time.Now()
	fmt.Printf("\n调度器开始运行... (任务总数: %d)\n", len(s.tasks))
	
	s.wg.Add(len(s.tasks))
	for i, task := range s.tasks {
		go s.worker(i+1, task, s.taskNames[i])
	}
	
	s.wg.Wait() // 等待所有任务完成
	totalTime := time.Since(s.startTime)
	
	fmt.Printf("\n所有任务已完成! 总耗时: %v\n", totalTime)
}

// ShowResults 显示任务执行结果
func (s *Scheduler) ShowResults() {
	fmt.Println("\n任务执行结果统计:")
	fmt.Println("---------------------------------")
	for _, result := range s.results {
		fmt.Printf("任务: %-15s 执行时间: %v\n", result.Name, result.Duration)
	}
	fmt.Println("---------------------------------")
}

func main() {
	// 创建任务调度器
	scheduler := NewScheduler()
	
	// 添加任务（名称和函数）
	scheduler.AddTask("数据清洗", func() {
		time.Sleep(time.Millisecond * 150)
	})
	
	scheduler.AddTask("文件处理", func() {
		time.Sleep(time.Millisecond * 200)
	})
	
	scheduler.AddTask("网络请求", func() {
		time.Sleep(time.Millisecond * 300)
	})
	
	scheduler.AddTask("图像处理", func() {
		time.Sleep(time.Millisecond * 100)
	})
	
	// 执行所有任务
	scheduler.Run()
	
	// 显示执行结果
	scheduler.ShowResults()
}
