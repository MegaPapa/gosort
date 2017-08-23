package sorts

import "math/rand"

func getMiddleElement(array []int) *int {
	currentSize := len(array)
	middleElement := array[int(currentSize / 2)]
	return &middleElement
}

func getLowAndHigh(array []int) (*int, *int) {
	return &array[0], &array[len(array) - 1]
}

/*func Quicksort(array []int) {
	if len(array) <= 2 {
		return
	}
	middleElement := getMiddleElement(array)
	low, high := getLowAndHigh(array)
	for element := range array {

	}
}*/

func Qsort(a []int) []int {
	if len(a) < 2 { return a }

	left, right := 0, len(a) - 1

	// Pick a pivot
	pivotIndex := rand.Int() % len(a)

	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	// Go down the rabbit hole
	Qsort(a[:left])
	Qsort(a[left + 1:])


	return a
}