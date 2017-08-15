package main

type sortResult struct {
	sortName string
	sortTime uint64
}

func swap(first, second *int) {
	var tmpPointer = *first
	*first = *second
	*second = tmpPointer
}