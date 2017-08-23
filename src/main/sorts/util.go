package sorts

type sortResult struct {
	sortName string
	sortTime uint64
}

func Swap(first, second *int) {
	var tmpPointer = *first
	*first = *second
	*second = tmpPointer
}