package main

import (
	"fmt"
	"sync"
)

/*
*

	编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。

*
*/
func send(sw *sync.WaitGroup, ch chan int) {
	// defer sw.Done()

	// for i := 1; i <= 10; i++ {
	// 	ch <- i
	// 	fmt.Printf("写入：", i)

	// }
	defer sw.Done()
	for i := 1; i <= 10; i++ {
		// ch <- i
		ch <- i
		fmt.Println("写入", i)
	}

}
func receive(sw *sync.WaitGroup, ch chan int) {
	// defer sw.Done()
	// for i := range ch {
	// 	fmt.Printf("取出：", i)
	// }
	defer sw.Done()
	for i := range ch {
		fmt.Println("接收", i)
	}
}
func main() {
	ch := make(chan int)

	var sw sync.WaitGroup
	sw.Add(2)

	go send(&sw, ch)
	go receive(&sw, ch)
	sw.Wait()

}
