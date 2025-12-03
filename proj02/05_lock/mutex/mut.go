package main

import (
	"fmt"
	"sync"
)

/*
*

	编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。

*
*/
func addTh(sw *sync.WaitGroup, lock *sync.Mutex, outnum *int) {
	defer sw.Done()

	for range 1000 {
		lock.Lock()
		*outnum++
		lock.Unlock()
	}

}
func main() {
	var sw sync.WaitGroup
	var lock sync.Mutex

	sw.Add(10)
	outnum := 0

	for range 10 {
		go addTh(&sw, &lock, &outnum)
	}

	sw.Wait()

	fmt.Printf("outnum:%d", outnum)
}
