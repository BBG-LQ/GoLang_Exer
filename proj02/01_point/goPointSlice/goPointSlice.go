package main

import (
	"fmt"
)

func pointSliceTwo(slicePoint *[]int) []int {
	for i, v := range *slicePoint {
		(*slicePoint)[i] *= 2
		fmt.Println(v)
	}
	return *slicePoint

}
func main() {
	pointSliceTwoNumbs := []int{2, 3, 4}
	fmt.Println(pointSliceTwo(&pointSliceTwoNumbs))

}
