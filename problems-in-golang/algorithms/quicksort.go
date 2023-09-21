package algorithms

import "fmt"

func TestQuickSort() {
	arr := []int{9, 3, 7, 4, 69, 420, 42}

	fmt.Println("Before quick sort: ", arr)
	qs(&arr, 0, len(arr)-1)
	fmt.Println("After quick sort: ", arr)
}

func qs(s *[]int, lo, hi int) {
	if lo < hi {
		pivotIdx := partition(s, lo, hi)
		qs(s, lo, pivotIdx-1)
		qs(s, pivotIdx+1, hi)
	}
}

// Partition takes a pivot and that pivot it divides both sides the lesser values will
// be on the left side and bigger values will be on the right side
func partition(s *[]int, lo, hi int) int {
	pivot := (*s)[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if (*s)[j] < pivot {
			(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
			i++
		}
	}
	(*s)[i], (*s)[hi] = (*s)[hi], (*s)[i]
	return i
}
