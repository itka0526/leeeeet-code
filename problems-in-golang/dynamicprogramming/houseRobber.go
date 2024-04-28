package dynamicprogramming

import (
	"github.com/itka0526/problems-in-golang/functions"
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
	n := len(nums) + 1
	dp := make([]int, n)
	dp[0] = 0
	dp[1] = nums[0]
	for i := 1; i < n-1; i++ {
		dp[i+1] = max(nums[i]+dp[i-1], dp[i])
	}
	return dp[len(dp)-1]
}

func TestRob() {
	s := []int{2, 7, 9, 3, 1}
	functions.LogAnswer(rob(s), 12)

	s1 := []int{2, 1, 1, 2}
	functions.LogAnswer(rob(s1), 4)

	s2 := []int{1, 3, 1}
	functions.LogAnswer(rob(s2), 3)
}
