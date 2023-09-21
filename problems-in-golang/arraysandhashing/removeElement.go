package arraysandhashing

import "fmt"

func removeElement(nums []int, val int) int {
	i := 0
	for _, num := range nums {
		if num != val {
			nums[i] = num
			i++
		}
	}
	return i
}

func TestRemoveElement() {
	s := []int{3, 2, 2, 3}

	l := removeElement(s, 3)

	fmt.Println(s[:l])
}
