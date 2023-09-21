package simple

import (
	"fmt"
)

func pivotInteger(n int) int {
	sum := n * (n + 1) / 2
	leftSum := 0
	for i := 1; i <= n; i++ {
		leftSum += i
		if sum-leftSum+i == leftSum {
			return i
		}
	}
	return -1
}

func TestPivotInteger() {
	fmt.Println("Answer: ", pivotInteger(8))
}
