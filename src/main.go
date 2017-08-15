package main

import "fmt"

func main() {
	a := 10
	b := 3
	swap(&a, &b);
	fmt.Printf("%d - %d", a, b)
	/*var matrix = []int {
		1, 2, 3,
		4, 10, 6,
		7, 8, 9,
	}
	var maxElement = matrix[0]
    for _, element := range matrix {
		if element > maxElement {
			maxElement = element
		}
	}
	fmt.Println(matrix)*/
}