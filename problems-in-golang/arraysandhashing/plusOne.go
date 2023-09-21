package arraysandhashing

import "fmt"

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+1 == 10 {
			digits[i] = 0
			if i-1 < 0 {
				digits = append([]int{0}, digits...)
				i++
			}
			continue
		}
		digits[i] = digits[i] + 1
		break
	}
	return digits
}

func TestPlusOne() {
	s := []int{9, 9}

	plusOne(s)

	fmt.Println(s)
}
