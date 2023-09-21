package twopointers

import (
	"fmt"
)

func maxArea(height []int) int {
	max := 0
	l, r := 0, len(height)-1

	for r > l {
		x := r - l
		y := min(height[l], height[r])

		s := x * y

		fmt.Println(s)

		if s > max {
			max = s
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestContainerWithMostWaterSolution() {
	// fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{2, 3, 4, 5, 18, 17, 6}))
}
