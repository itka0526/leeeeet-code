package dynamicprogramming

import (
	"fmt"
	"strconv"
)

func numDecodings(s string) int {
	dp := map[int]int{len(s): 1}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			dp[i] = 0
		} else {
			dp[i] = dp[i+1]
		}
		if i+1 < len(s) {
			if n, err := strconv.Atoi(s[i : i+1+1]); err == nil && n >= 10 && n <= 26 {
				fmt.Println(n)
				dp[i] += dp[i+2]
			}
		}
	}
	fmt.Println(dp)
	return dp[0]
}
