package arraysandhashing

import "fmt"

func topKFrequentSolution(nums []int, k int) (res []int) {
	// slice to map
	m := make(map[int]int)
	r := make([][]int, len(nums)+1)

	for i := range nums {
		m[nums[i]] += 1
	}

	for key := range m {
		r[m[key]] = append(r[m[key]], key)
	}

	for i := len(r) - 1; i > 0; i-- {
		res = append(res, r[i]...)

		if len(res) == k {
			return
		}
	}

	return
}

func TestTopKFrequent() {
	nums := []int{1, 1, 1, 2, 2, 3}
	nums2 := []int{-1, -1}

	k := 2
	k2 := 1

	fmt.Println(topKFrequentSolution(nums, k))
	fmt.Println(topKFrequentSolution(nums2, k2))
}
