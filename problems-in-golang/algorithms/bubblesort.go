package algorithms

import "fmt"

func bubbleSort(s *[]int) {
	for i := 0; i < len(*s); i++ {
		for j := 0; j < len(*s)-1-i; j++ {
			if (*s)[j] > (*s)[j+1] {
				(*s)[j], (*s)[j+1] = (*s)[j+1], (*s)[j]
			}
		}
	}
}

func TestBubbleSort() {
	array := []int{11, 14, 3, 8, 18, 17, 43}

	fmt.Println("Before bubble sort: ", array)
	bubbleSort(&array)
	fmt.Println("After bubble sort: ", array)
}
