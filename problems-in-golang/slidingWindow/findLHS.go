package slidingwindow

import (
	"fmt"
	"math"
)

func findLHS(nums []int) int {
	m := map[int]int{}
	for _, n := range nums {
		m[n]++
	}
	max := math.MinInt32
	for k, v := range m {
		if v2, ok := m[k+1]; ok {
			if max < v+v2 {
				max = v + v2
				continue
			}
		}
	}
	if max == math.MinInt32 {
		return 0
	}
	return max
}

func TestFindLHS() {
	fmt.Println("Expt: 5; Got: ", findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}))
}
