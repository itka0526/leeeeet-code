package dynamicprogramming

import (
	"fmt"
)

func climbStairs1(n int) int {
	o, t := 1, 1
	for i := 0; i < n; i++ {
		tmp := o
		o = o + t
		t = tmp
	}
	return o
}

// Time Limit Exceeds
func climbStairs2(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs2(n-1) + climbStairs2(n-2)
}

var mem = map[int]int{}

func climbStairsRecursive(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else if val, ok := mem[n]; ok {
		return val
	}
	result := climbStairsRecursive(n-1) + climbStairsRecursive(n-2)
	mem[n] = result
	return result
}

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[n] = 1
	dp[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		dp[i] = dp[i+1] + dp[i+2]
	}
	return dp[0]
}

func TestClimbStair() {
	fmt.Println("Possibilities: ", climbStairs(4))
	fmt.Println("Possibilities: ", climbStairs(5))
}
