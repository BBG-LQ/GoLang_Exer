package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
*

	使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。

*
*/
func addTh(sw *sync.WaitGroup, outnum *atomic.Int32) {
	defer sw.Done()
	for range 1000 {
		outnum.Add(1)
	}
}
func main() {
	var sw sync.WaitGroup
	sw.Add(10)
	var outnum atomic.Int32

	for i := range 10 {
		go addTh(&sw, &outnum)
		fmt.Printf("work %d\n", i)
	}

	sw.Wait()
	fmt.Printf("outnum:%d", outnum)
}
