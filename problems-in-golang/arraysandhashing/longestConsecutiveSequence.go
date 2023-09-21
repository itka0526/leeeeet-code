package arraysandhashing

import (
	"fmt"
	"runtime/debug"
)

func longestConsecutiveSolution(nums []int) int {
	debug.SetGCPercent(1)

	m := make(map[int]bool)

	for idx := range nums {
		m[nums[idx]] = true
	}

	max := 0

	for num := range m {
		if m[num-1] {
			continue
		}

		l := 1
		next := num + 1

		for m[next] {
			l++
			next++
		}

		if l > max {
			max = l
		}
	}

	return max
}

func TestLongestConsecutive() {
	n1 := []int{100, 4, 200, 1, 3, 2}

	fmt.Println("Solution: ", longestConsecutiveSolution(n1))

	n2 := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println("Solution: ", longestConsecutiveSolution(n2))

	n3 := []int{0}
	fmt.Println("Failing test: ", longestConsecutiveSolution(n3))

	n4 := []int{}
	fmt.Println("Failing test: ", longestConsecutiveSolution(n4))
}
