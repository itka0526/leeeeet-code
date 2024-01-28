package seventyfive

import "sort"

func maxOperations(nums []int, k int) int {
	res, l, r := 0, 0, len(nums)-1
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for l < r {
		s := nums[l] + nums[r]
		if s == k {
			res++
			l++
			r--
		} else if s > k {
			r--
		} else {
			l++
		}
	}
	return res
}
