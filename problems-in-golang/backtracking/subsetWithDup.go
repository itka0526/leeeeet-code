package backtracking

import "sort"

func subsetsWithDup(nums []int) (res [][]int) {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	var dfs func(int, []int)
	dfs = func(i int, s []int) {
		if i == len(nums) {
			cp := make([]int, len(s))
			copy(cp, s)
			res = append(res, cp)	
			return
		}
		s = append(s, nums[i])
		dfs(i+1, s)
		s = s[:len(s)-1]
		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}
		dfs(i+1, s)
	}

	dfs(0, []int{})
	return
}

func TestSubsetWithDup() {
	subsetsWithDup([]int{1, 2, 2})
}
