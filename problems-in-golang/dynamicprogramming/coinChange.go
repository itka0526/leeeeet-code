package dynamicprogramming

// dfs approach?
func coinChangeSolution1(coins []int, amount int) int {
	memo := map[int]int{}
	var dfs func(sum int) int
	dfs = func(sum int) int {
		if sum == amount {
			return 0
		}
		if val, ok := memo[sum]; ok {
			return val
		}
		length := -1
		for _, coin := range coins {
			if amount-sum >= coin {
				opt := dfs(sum + coin)
				if opt >= 0 {
					if length == -1 || length > opt+1 {
						length = opt + 1
					}
				}
			}
		}
		memo[sum] = length
		return length
	}
	return dfs(0)
}

// dp
func coinChangeSolution2(coins []int, amount int) int {
	impossibleAmount := amount + 1
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = impossibleAmount
	}
	dp[0] = 0
	for val := 1; val < amount+1; val++ {
		for _, c := range coins {
			if c <= val {
				dp[val] = min(dp[val], 1+dp[val-c])
			}
		}
	}
	if dp[amount] == impossibleAmount {
		return -1
	}
	return dp[amount]
}

func coinChange(a []int, k int) int {
	b := make([]int, k+1)
	for i := range b {
		b[i] = k + 1
	}
	b[0] = 0
	for i := 1; i < len(b); i++ {
		for _, c := range a {
			if i >= c {
				b[i] = min(b[i], 1+b[i-c])
			}
		}
	}
	if b[k] == k+1 {
		return -1
	}
	return b[k]
}

func cp(a []int) (b []int) {
	b = make([]int, len(a))
	copy(b, a)
	return b
}
