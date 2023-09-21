package arraysandhashing

import "fmt"

func twoSumSolution(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if k, ok := m[target-nums[i]]; ok {
			return []int{i, k}
		}
		m[nums[i]] = i
	}

	return []int{}
}

func TestTwoSum() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSumSolution(nums, target))
}
