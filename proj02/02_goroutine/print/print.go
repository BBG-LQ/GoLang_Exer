package main

import (
	"fmt"
	"sync"
)

func printOdd1(sw *sync.WaitGroup) {
	defer sw.Done()
	for i := 1; i <= 10; i += 2 {
		fmt.Println("odd:", i)
	}

}
func printeven1(sw *sync.WaitGroup) {
	defer sw.Done()
	for i := 2; i <= 10; i += 2 {
		fmt.Println("even:", i)
	}
}
func main() {
	var sw sync.WaitGroup
	sw.Add(2)

	go printOdd1(&sw)
	go printeven1(&sw)

	sw.Wait()

}
