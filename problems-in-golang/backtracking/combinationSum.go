package backtracking

import (
	"fmt"
)

func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	subset := make([]int, 0)

	var dfs func(int, int)

	dfs = func(idx int, total int) {
		if total == target {
			cp := make([]int, len(subset))
			copy(cp, subset)
			ans = append(ans, cp)
			return
		}

		if idx >= len(candidates) || total > target {
			return
		}

		subset = append(subset, candidates[idx])
		dfs(idx, total+candidates[idx])

		subset = subset[:len(subset)-1]
		dfs(idx+1, total)
	}

	dfs(0, 0)

	return ans
}

func TestCombinationSum() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
