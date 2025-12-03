package main

import (
	"fmt"
	"sync"
)

/**

	实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。

**/

func send(sw *sync.WaitGroup, ch chan int) {
	defer sw.Done()
	for i := 1; i <= 100; i++ {
		ch <- i
		fmt.Println("写入", i)
	}

}
func receive(sw *sync.WaitGroup, ch chan int) {
	defer sw.Done()
	for i := 1; i <= 100; i++ {
		num := <-ch
		fmt.Println("re:", num)
	}

}
func main() {
	ch := make(chan int, 100)
	var sw sync.WaitGroup

	sw.Add(2)
	go send(&sw, ch)

	go receive(&sw, ch)

	sw.Wait()

}
