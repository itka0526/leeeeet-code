package seventyfive

import (
	"fmt"
)

func findMaxAverage(nums []int, k int) float64 {
	l, r := 0, 0
	max := 0.0
	sum := 0.0
	for r < k {
		sum += float64(nums[r])
		r++
	}
	max = sum
	for r < len(nums) {
		// Add one more element but remove the first element
		// So the length stays the same
		// This approach removes the unnecessary recalculations
		sum += float64(nums[r]) - float64(nums[l])
		if max <= sum {
			max = sum
		}
		l++
		r++
	}
	return max / float64(k)
}

func TestFindMaxAverage() {
	fmt.Println("Exp: -1.00000; Got: ", findMaxAverage([]int{-1}, 1))
	fmt.Println("Exp: 5; Got: ", findMaxAverage([]int{5}, 1))
}
