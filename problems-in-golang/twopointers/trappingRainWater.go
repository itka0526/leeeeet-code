package twopointers

import (
	"fmt"
)

// first implementation
func trap1(height []int) int {
	left := make([]int, len(height))
	right := make([]int, len(height))

	leftMax, rightMax := height[0], height[len(height)-1]
	l, r := 0, len(height)-1

	for range height {
		if height[l] > leftMax {
			leftMax = height[l]
		}

		if height[r] > rightMax {
			rightMax = height[r]
		}

		left[l] = leftMax
		right[r] = rightMax

		l++
		r--
	}

	water := 0

	for i := range height {
		wall := minInt(left[i], right[i])
		possibleWater := wall - height[i]

		if possibleWater > 0 {
			water += possibleWater
		}
	}

	return water
}

func trap(height []int) int {
	// 0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1
	if height == nil {
		return 0
	}

	// and we want update left and right baseid min values
	l, r, w := 0, len(height)-1, 0
	leftMax, rightMax := height[l], height[r]

	for r > l {
		if rightMax > leftMax {
			l++
			leftMax = maxInt(leftMax, height[l])

			w += leftMax - height[l]
		} else {
			r--
			rightMax = maxInt(rightMax, height[r])

			w += rightMax - height[r]
		}
	}

	return w
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestTrap() {
	nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	fmt.Println("Solution: ", trap(nums))
}
