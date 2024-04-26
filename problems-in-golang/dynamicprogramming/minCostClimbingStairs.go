package dynamicprogramming

import "fmt"

func minCostClimbingStairsTabulated(cost []int) int {
	dp := []int{}
	for range cost {
		dp = append(dp, 0)
	}
	dp = append(dp, 0)
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = cost[i] + min(dp[i-2], dp[i-1])
	}
	return min(dp[len(cost)-1], dp[len(cost)-2])
}

func minCostClimbingStairs(cost []int) int {
	f, s := cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		c := cost[i] + min(f, s)
		f = s
		s = c
	}
	return min(f, s)
}

func TestMinCostClimbingStairs() {
	s := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println("Answer: ", minCostClimbingStairs(s), "; Expected: ", 6)
}
