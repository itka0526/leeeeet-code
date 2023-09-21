package arraysandhashing

import (
	"fmt"
)

func productExceptionSelfSolution(nums []int) []int {
	a := make([]int, len(nums))

	// prefix
	for i, m := 0, 1; i < len(nums); i++ {
		a[i], m = m, m*nums[i]
	}

	b := make([]int, len(nums))
	// postfix
	for i, m := len(nums)-1, 1; i >= 0; i-- {
		b[i], m = a[i]*m, m*nums[i]
	}

	return b
}

func TestProductExceptionSelfSolution() {
	fmt.Println(productExceptionSelfSolution([]int{1, 2, 3, 4}))
	// fmt.Println(productExceptionSelfSolution([]int{0, 0}))
	// fmt.Println(productExceptionSelfSolution([]int{-1, 1, 0, -3, 3}))
	// fmt.Println(productExceptionSelfSolution([]int{1, -1}))
}
