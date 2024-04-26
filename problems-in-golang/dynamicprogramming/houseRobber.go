package dynamicprogramming

import (
	"fmt"
	"math"
)

// naive approach
// func rob(nums []int) int {
// 	a, b := 0, 0
// 	for i := range nums {
// 		if i%2 == 0 {
// 			a += nums[i]
// 		} else {
// 			b += nums[i]
// 		}
// 	}
// 	return max(a, b)
// }

func rec(nums []int, start int, currSum int, path int, maxSum *int) int {
	if start > len(nums)-1 {
		if path > *maxSum {
			*maxSum = path
		}
		return 0
	}
	path += nums[start]
	currSum += nums[start]
	for i := start; i < len(nums); i++ {
		call := rec(nums, i+2, currSum, path, maxSum)
		currSum = max(call, currSum)
	}
	return currSum
}

func rob(nums []int) int {
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	maxSum := math.MinInt
	rec(nums, 0, 0, 0, &maxSum)
	return maxSum
}

func LogAnswer[T any](
	got T,
	exp T,
) {
	if _, err := fmt.Println("Answer: ", got, "; Exp: ", exp); err != nil {
		panic("Cannot log this!")
	}
}

func TestRob() {
	s := []int{2, 7, 9, 3, 1}
	LogAnswer(rob(s), 12)

	s1 := []int{2, 1, 1, 2}
	LogAnswer(rob(s1), 4)

	s2 := []int{1, 3, 1}
	LogAnswer(rob(s2), 3)
}
