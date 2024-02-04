package seventyfive

import (
	"fmt"
)

func longestSubarray(nums []int) int {
	l, r := 0, 0
	m := 0
	s := 1
	for r < len(nums) {
		if nums[r] == 0 {
			s--
		}
		for s < 0 {
			if nums[l] == 0 {
				s++
			}
			l++
		}
		r++
		if r-l-1 > m {
			m = r - l - 1
		}
	}
	return m
}

func TestLongestSubarray() {
	fmt.Println(longestSubarray([]int{1, 1, 0, 1}))
	fmt.Println("--------------------------------")
	fmt.Println(longestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1}))
}
