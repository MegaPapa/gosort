package main

import (
	"fmt"
)

func main() {
	array := [9]int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(array)
	Arr(array[:])
	fmt.Println(array)
}

func Arr(array []int) {
	for i := 0; i < len(array); i++ {
		array[i]++
	}
}