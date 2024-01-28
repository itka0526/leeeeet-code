package backtracking

import (
	"fmt"
)

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	subset := make([]int, 0)

	var dfs func(int)
	dfs = func(i int) {
		if i >= len(nums) {
			cp := make([]int, len(subset))
			copy(cp, subset)

			ans = append(ans, cp)
			return
		}

		subset = append(subset, nums[i])
		dfs(i + 1)

		subset = subset[:len(subset)-1]
		dfs(i + 1)
	}

	dfs(0)
	return ans
}

func TestSubsets() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
