package arraysandhashing

import "fmt"

func containsDuplicateSolution(nums []int) bool {
	// key = 'element', value = 'count'
	m := map[int]int{}

	for i := 0; i < len(nums); i++ {
		if m[nums[i]] == 1 {
			return false
		}
		m[nums[i]] += 1
	}

	return true
}

func TestContainsDuplicate() {
	nums := []int{1, 2, 3, 1}
	nums2 := []int{1, 2, 3, 4}
	nums3 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}

	fmt.Println(containsDuplicateSolution(nums))
	fmt.Println(containsDuplicateSolution(nums2))
	fmt.Println(containsDuplicateSolution(nums3))
}
