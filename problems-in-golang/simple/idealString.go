package simple

import (
	"fmt"
	"math"
)

func longestIdealString(s string, k int) int {
	dp := make([]int, 27)
	n := len(s)
	for i := n - 1; i >= 0; i-- {
		c := s[i]
		idx := int(c - 'a')
		maxi := math.MinInt
		l, r := max(idx-k, 0), min(idx+k, 26)
		for j := l; j <= r; j++ {
			maxi = max(maxi, dp[j])
		}
		dp[idx] = maxi + 1
	}
	m := math.MinInt
	for _, val := range dp {
		if val > m {
			m = val
		}
	}
	return m
}

func TestLongestIdealString() {
	s := "acfgbd"
	fmt.Println(longestIdealString(s, 2))
}
