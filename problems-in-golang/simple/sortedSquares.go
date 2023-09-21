package simple

import "sort"

func sortedSquares(nums []int) []int {
	s := make([]int, len(nums))
	for idx, i := range nums {
		s[idx] = i * i
	}
	sort.Slice(s, func(a, b int) bool { return s[a] < s[b] })
	return s
}
