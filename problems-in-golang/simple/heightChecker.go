package simple

import "sort"

func heightChecker(heights []int) int {
	exp := make([]int, len(heights))

	copy(exp, heights)
	sort.Slice(exp, func(i, j int) bool { return exp[i] < exp[j] })

	var unexp int
	for i := 0; i < len(heights); i++ {
		if exp[i] != heights[i] {
			unexp++
		}
	}
	return unexp
}
