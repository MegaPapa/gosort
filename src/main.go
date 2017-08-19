package main

import (
	"fmt"
	"./sorts"
)

func main() {
	var array [10]int
	fmt.Println(array)
	arr(array[:])
	fmt.Println(array)
	quicksort(array) 
}

func arr(array []int) {
	for i := 0; i < len(array); i++ {
		array[i]++
	}
}