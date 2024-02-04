package arraysandhashing

import "fmt"

func shuffle(nums []int, n int) []int {
	r := make([]int, len(nums))
	for i := 0; i < n; i++ {
		r = append(r, nums[i], nums[i+n])
	}
	return r
}

func TestShuffle() {
	fmt.Println("Exp: [2 3 5 4 1 7]; Got: ", shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
}
