package arraysandhashing

import "fmt"

func removeDuplicates(nums []int) int {
	// key - element, value - index
	return helperFunction(&nums)
}

func helperFunction(nums *[]int) int {
	m := make(map[int]int)
	idx := 0
	for i := range *nums {
		if _, ok := m[(*nums)[i]]; !ok {
			m[(*nums)[i]] = idx
			idx++
		}
	}
	(*nums) = (*nums)[:len(m)]
	for el, idx := range m {
		(*nums)[idx] = el
	}
	return len(*nums)
}

func removeDuplicates1(nums []int) int {
	// two pointer solution?
	i := 0
	j := i + 1

	for j < len(nums) {
		if (nums)[i] == (nums)[j] {
			nums = append((nums)[:i], (nums)[j:]...)
			continue
		}
		i++
		j++
	}
	return len(nums)
}

func perfomantRemoveDuplicate(nums []int) int {
	index := 0
	i, j := 0, 0

	for i < len(nums) {
		j = i
		for j < len(nums) && nums[j] == nums[i] {
			j++
		}
		nums[index] = nums[i]
		index++
		i = j
	}

	return index
}

func TestRemoveDuplicates() {
	s := []int{1, 1, 2}

	c := perfomantRemoveDuplicate(s)
	fmt.Println(c, s)
}
