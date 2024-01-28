package arraysandhashing

import "fmt"

func MoveZeroesToEnd() {
	a := []int{0, 1, 0, 5, 12, 2, 4, 5, 0, 1, 2, 3, 0, 44, 0, 55, 0, 0, 0, 0, 1}

	fmt.Println("Before: ", a)
	l, r := 0, 0

	for r < len(a) {
		// switch with the first non-zero value
		// l keep track of zero
		// r find the first non-zero value
		if a[r] != 0 {
			a[r], a[l] = a[l], a[r]
			l++
		}
		if a[r] == 0 && a[l] != 0 {
			l = r
		}
		r++
	}

	fmt.Println("After: ", a)
}
