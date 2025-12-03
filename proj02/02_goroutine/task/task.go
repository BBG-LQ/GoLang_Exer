package main

import (
	"fmt"
	"sync"
	"time"
)

type Task func()

// 专用协程启动函数
func startTask(wg *sync.WaitGroup, index int, task Task) {
	defer wg.Done()

	// 时间统计逻辑
	start := time.Now()
	fmt.Printf("任务 %d 执行开始时间: %v\n", index+1, start)
	defer func() {
		elapsed := time.Since(start)
		fmt.Printf("任务 %d 执行结束时间: %v\n", index+1, time.Now())
		fmt.Printf("任务 %d 执行用时: %v\n", index+1, elapsed)
	}()

	// 执行任务
	task()
}

func runTasks(Tasks []Task) {
	var sw sync.WaitGroup
	for i, task := range Tasks {
		sw.Add(1)

		go startTask(&sw, i, task)

	}
	sw.Wait()
}

func main() {

	Tasks := []Task{
		func() {
			time.Sleep(1 * time.Second)
			fmt.Println("first")
		},
		func() {
			time.Sleep(2 * time.Second)
			fmt.Println("second")
		},
		func() {
			time.Sleep(3 * time.Second)
			fmt.Println("third")
		},
	}

	runTasks(Tasks)
}
